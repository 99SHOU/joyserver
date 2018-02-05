package internal

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
)

type NodeServerHandler struct {
	Node      *Node
	processor *net.Processor
}

func (h *NodeServerHandler) Register(server *net.Server) {
	server.OnNewAgent = h.NewAgent
	server.OnCloseAgent = h.CloseAgent

	h.processor = server.Processor
	msg.RegisterMsg(server.Processor)
	server.SetHandler(pb.MsgID_FWM_NODE_REGISTER_REQ, h.OnNodeRegisterReq)
	server.SetHandler(pb.MsgID_FWM_GAME_MSG_TRANSFER, h.OnGameMsgTranfer)
}

func (h *NodeServerHandler) NewAgent(agent net.Agent) {
}

func (h *NodeServerHandler) CloseAgent(agent net.Agent) {
}

func (h *NodeServerHandler) OnNodeRegisterReq(message interface{}, agent interface{}) {
	msg := message.(*pb.NodeRegisterReq)
	a := agent.(*net.BaseAgent)

	if msg.NodeType == pb.NODE_TYPE_GATE {
		a.SetNodeType(msg.NodeType)
		a.SetNodeStatu(msg.NodeStatu)
		a.SetNodeID(define.NodeID(msg.NodeId))
		a.SetAgentInfo(pb.AGENT_INFO_KEY_NODE_PORT, msg.NodePort)

		log.Debug("New node register to game NodeType: %v NodeID: %v", msg.NodeType.String(), msg.NodeId)

		a.WriteMsg(&pb.NodeRegisterAck{NodeType: h.Node.NodeType, NodeId: uint32(h.Node.NodeID)})
	}
}

func (h *NodeServerHandler) OnGameMsgTranfer(message interface{}, agent interface{}) {
	msg := message.(*pb.GameMsgTransfer)
	data := make([]byte, len(msg.MsgId)+len(msg.MsgBody))
	copy(data, msg.MsgId)
	copy(data[len(msg.MsgId):], msg.MsgBody)

	buildinMsg, err := h.processor.Unmarshal(data)
	if err != nil {
		log.Debug("unmarshal message error: %v  %v", err, data)
		return
	}

	h.processor.Dispatch(buildinMsg, agent)
}
