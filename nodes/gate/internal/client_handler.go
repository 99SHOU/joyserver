package internal

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/name5566/leaf/log"
)

type ClientHandler struct {
	Node *Node
}

func (ch *ClientHandler) Register(client *net.Client) {
	client.OnNewAgent = ch.NewAgent
	client.OnCloseAgent = ch.CloseAgent

	client.SetHandler(pb.MsgID_FWM_NODE_REGISTER_ACK, ch.OnNodeRegisterAck)
}

func (ch *ClientHandler) NewAgent(agent net.Agent) {
	ch.Node.AgentManager.AddAgent(agent.RemoteAddr().String(), agent)

	ch.NodeRegisterReq(agent)
}

func (ch *ClientHandler) CloseAgent(agent net.Agent) {
	ch.Node.AgentManager.RemoveAgent(agent.RemoteAddr().String())
}

func (ch *ClientHandler) NodeRegisterReq(agent net.Agent) {
	agent.WriteMsg(&pb.NodeRegisterReq{NodeType: ch.Node.NodeType, NodeId: uint32(ch.Node.NodeID)})
}

func (ch *ClientHandler) OnNodeRegisterAck(message interface{}, agent interface{}) {
	msg := message.(*pb.NodeRegisterAck)
	a := agent.(*net.BaseAgent)
	a.SetNodeType(msg.NodeType)
	a.SetNodeID(define.NodeID(msg.NodeId))

	log.Debug("Register to NodeType: %v NodeID: %v success", msg.NodeType.String(), msg.NodeId)
}
