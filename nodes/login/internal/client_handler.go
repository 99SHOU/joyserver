package internal

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/msg"
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

	msg.RegisterMsg(client.Processor)
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
	agent.WriteMsg(&pb.NodeRegisterReq{NodeType: ch.Node.NodeType, NodeStatu: ch.Node.NodeStatu, NodeId: uint32(ch.Node.NodeID), ServerPort: ch.Node.NodeCfg.Port})
}

func (ch *ClientHandler) OnNodeRegisterAck(message interface{}, agent interface{}) {
	msg := message.(*pb.NodeRegisterAck)
	a := agent.(*net.BaseAgent)
	a.SetNodeType(msg.NodeType)
	a.SetNodeStatu(msg.NodeStatu)
	a.SetNodeID(define.NodeID(msg.NodeId))

	if a.GetNodeType() == pb.NODE_TYPE_CENTER {
		ch.Node.NodeStatu = pb.NODE_STATU_READY
		allAgent := ch.Node.AgentManager.GetAgentAll()
		net.BroadcastMsg(allAgent, &pb.SetNodeStatu{NodeStatu: ch.Node.NodeStatu})
	}

	log.Debug("Register to NodeType: %v NodeID: %v success", msg.NodeType.String(), msg.NodeId)
}
