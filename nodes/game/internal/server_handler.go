package internal

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
)

type ServerHandler struct {
	Node *Node
}

func (sh *ServerHandler) Register(server *net.Server) {
	server.OnNewAgent = sh.NewAgent
	server.OnCloseAgent = sh.CloseAgent

	msg.RegisterMsg(server.Processor)
	server.SetHandler(pb.MsgID_FWM_NODE_REGISTER_REQ, sh.OnNodeRegisterReq)
}

func (sh *ServerHandler) NewAgent(agent net.Agent) {
}

func (sh *ServerHandler) CloseAgent(agent net.Agent) {
}

func (sh *ServerHandler) OnNodeRegisterReq(message interface{}, agent interface{}) {
	msg := message.(*pb.NodeRegisterReq)
	a := agent.(*net.BaseAgent)

	if msg.NodeType == pb.NODE_TYPE_GATE {
		a.SetNodeType(msg.NodeType)
		a.SetNodeStatu(msg.NodeStatu)
		a.SetNodeID(define.NodeID(msg.NodeId))
		a.SetAgentInfo(pb.NODE_INFO_SERVER_PORT, msg.ServerPort)

		log.Debug("New node register to center NodeType: %v NodeID: %v", msg.NodeType.String(), msg.NodeId)

		a.WriteMsg(&pb.NodeRegisterAck{NodeType: sh.Node.NodeType, NodeId: uint32(sh.Node.NodeID)})
	}
}
