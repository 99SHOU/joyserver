package internal

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
)

type NodeClientHandler struct {
	Node      *Node
	processor *net.Processor
}

func (h *NodeClientHandler) Register(client *net.Client) {
	client.OnNewAgent = h.NewAgent
	client.OnCloseAgent = h.CloseAgent

	h.processor = client.Processor
	msg.RegisterMsg(client.Processor)
	client.SetHandler(pb.MsgID_FWM_NODE_REGISTER_ACK, h.OnNodeRegisterAck)
	client.SetHandler(pb.MsgID_FWM_GAME_NODE_LIST_ACK, h.OnGameNodeListAck)
	client.SetHandler(pb.MsgID_FWM_GAME_MSG_TRANSFER, h.OnGameMsgTransfer)
}

func (h *NodeClientHandler) NewAgent(agent net.Agent) {
	h.Node.AgentManager.AddAgent(agent.RemoteAddr().String(), agent)

	h.NodeRegisterReq(agent)
}

func (h *NodeClientHandler) CloseAgent(agent net.Agent) {
	h.Node.AgentManager.RemoveAgent(agent.RemoteAddr().String())
}

func (h *NodeClientHandler) NodeRegisterReq(agent net.Agent) {
	agent.WriteMsg(&pb.NodeRegisterReq{NodeType: h.Node.NodeType, NodeStatu: h.Node.NodeStatu, NodeId: uint32(h.Node.NodeID), NodePort: h.Node.NodeCfg.NodePort})
}

func (h *NodeClientHandler) OnNodeRegisterAck(message interface{}, agent interface{}) {
	msg := message.(*pb.NodeRegisterAck)
	a := agent.(*net.BaseAgent)
	a.SetNodeType(msg.NodeType)
	a.SetNodeStatu(msg.NodeStatu)
	a.SetNodeID(define.NodeID(msg.NodeId))

	if a.GetNodeType() == pb.NODE_TYPE_CENTER {
		h.Node.NodeStatu = pb.NODE_STATU_READY
		allAgent := h.Node.AgentManager.GetAgentAll()
		net.BroadcastMsg(allAgent, &pb.SetNodeStatu{NodeStatu: h.Node.NodeStatu})
		a.WriteMsg(&pb.GameNodeListReq{})
	}

	log.Debug("Register to NodeType: %v NodeID: %v success", msg.NodeType.String(), msg.NodeId)
}

func (h *NodeClientHandler) OnGameNodeListAck(message interface{}, agent interface{}) {
	msg := message.(*pb.GameNodeListAck)

	for _, agentInfo := range msg.NodeInfos {
		if !h.Node.ClientManager.ExistClient(agentInfo.Addr) {
			log.Debug("Connect to game NodeId: %v, NodeType: %v, NodeStatu: %v", agentInfo.GetNodeId(), agentInfo.GetNodeType(), agentInfo.GetNodeStatu())

			h.Node.ClientManager.NewAndStart(agentInfo.Addr, &NodeClientHandler{Node: h.Node}, net.NewProcessor())
		}
	}
}

func (h *NodeClientHandler) OnGameMsgTransfer(message interface{}, agent interface{}) {
	msg := message.(*pb.GameMsgTransfer)
	data := make([]byte, len(msg.MsgId)+len(msg.MsgBody))
	copy(data, msg.MsgId)
	copy(data[len(msg.MsgId):], msg.MsgBody)

	buildinMsg, err := h.processor.Unmarshal(data)
	if err != nil {
		log.Debug("unmarshal message error: %v  %v", err, data)
		return
	}

	characterId := msg.CharacterId
	a := h.Node.AgentManager.GetAgentByNodeInfo(pb.AGENT_INFO_KEY_CHARACTER_ID, characterId)
	if a != nil && len(a) > 0 {
		a[0].WriteMsg(buildinMsg)
	}

	//sh.processor.Dispatch(buildinMsg, agent)
}
