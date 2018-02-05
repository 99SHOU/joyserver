package internal

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
	"strconv"
	"strings"
)

type NodeServerHandler struct {
	Node *Node
}

func (h *NodeServerHandler) Register(server *net.Server) {
	server.OnNewAgent = h.NewAgent
	server.OnCloseAgent = h.CloseAgent

	msg.RegisterMsg(server.Processor)
	server.SetHandler(pb.MsgID_FWM_NODE_REGISTER_REQ, h.OnNodeRegisterReq)
	server.SetHandler(pb.MsgID_FWM_SET_NODE_STATU, h.OnSetNodeStatu)
	server.SetHandler(pb.MsgID_FWM_GAME_NODE_LIST_REQ, h.OnGameNodeListReq)
}

func (h *NodeServerHandler) NewAgent(agent net.Agent) {
	h.Node.AgentManager.AddAgent(agent.RemoteAddr().String(), agent)
}

func (h *NodeServerHandler) CloseAgent(agent net.Agent) {
	h.Node.AgentManager.RemoveAgent(agent.RemoteAddr().String())
}

func (h *NodeServerHandler) OnNodeRegisterReq(message interface{}, agent interface{}) {
	msg := message.(*pb.NodeRegisterReq)
	a := agent.(*net.BaseAgent)

	a.SetNodeType(msg.NodeType)
	a.SetNodeStatu(pb.NODE_STATU_NOT_READY)
	a.SetNodeID(define.NodeID(msg.NodeId))
	a.SetAgentInfo(pb.AGENT_INFO_KEY_NODE_PORT, msg.NodePort)

	log.Debug("New node register to center NodeType: %v NodeID: %v", msg.NodeType.String(), msg.NodeId)

	a.WriteMsg(&pb.NodeRegisterAck{NodeType: h.Node.NodeType, NodeId: uint32(h.Node.NodeID)})
}

func (h *NodeServerHandler) OnSetNodeStatu(message interface{}, agent interface{}) {
	msg := message.(*pb.SetNodeStatu)
	a := agent.(*net.BaseAgent)

	a.SetNodeStatu(msg.NodeStatu)

	log.Debug("Change node statu : NodeId: %v NodeType: %v NodeStatu: %v", a.GetNodeID(), a.GetNodeType(), a.GetNodeStatu())

	if a.GetNodeType() == pb.NODE_TYPE_GAME {
		gateAgent := h.Node.AgentManager.GetAgentByNodeType([]pb.NODE_TYPE{pb.NODE_TYPE_GATE})
		net.BroadcastMsg(gateAgent, h.BuildGameNodeList())
	}
}

func (h *NodeServerHandler) OnGameNodeListReq(message interface{}, agent interface{}) {
	a := agent.(*net.BaseAgent)

	a.WriteMsg(h.BuildGameNodeList())
}

func (h *NodeServerHandler) BuildGameNodeList() *pb.GameNodeListAck {
	gameAgents := h.Node.AgentManager.GetAgentByNodeType([]pb.NODE_TYPE{pb.NODE_TYPE_GAME})
	gameNodeList := pb.GameNodeListAck{}
	gameNodeList.NodeInfos = []*pb.NodeInfo{}

	for _, agent := range gameAgents {
		if agent.GetNodeStatu() == pb.NODE_STATU_READY {

			nodePort := agent.GetAgentInfo(pb.AGENT_INFO_KEY_NODE_PORT).(uint32)
			nodeAddr := strings.Split(agent.RemoteAddr().String(), ":")[0] + ":" + strconv.FormatUint(uint64(nodePort), 10)

			gameNodeList.NodeInfos = append(gameNodeList.NodeInfos, &pb.NodeInfo{NodeType: agent.GetNodeType(), NodeStatu: agent.GetNodeStatu(), NodeId: uint32(agent.GetNodeID()), Addr: nodeAddr})
		}
	}

	return &gameNodeList
}
