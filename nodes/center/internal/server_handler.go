package internal

import (
	"github.com/99SHOU/joyserver/common/define"
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

	server.SetHandler(pb.MsgID_FWM_NODE_REGISTER_REQ, sh.OnNodeRegisterReq)
}

func (sh *ServerHandler) NewAgent(agent *net.ServerAgent) {
	sh.Node.AgentManager.AddAgent(agent.RemoteAddr().String(), agent)
}

func (sh *ServerHandler) CloseAgent(agent *net.ServerAgent) {
	sh.Node.AgentManager.RemoveAgent(agent.RemoteAddr().String())
}

func (sh *ServerHandler) OnNodeRegisterReq(message interface{}, agent interface{}) {
	msg := message.(*pb.NodeRegisterReq)
	a := agent.(*net.ServerAgent)

	a.SetNodeType(msg.NodeType)
	a.SetNodeID(define.NodeID(msg.NodeId))

	log.Debug("New node register to center NodeType: %v NodeID: %v", msg.NodeType.String(), msg.NodeId)

	a.WriteMsg(&pb.NodeRegisterAck{NodeType: sh.Node.NodeType, NodeId: uint32(sh.Node.NodeID)})
}
