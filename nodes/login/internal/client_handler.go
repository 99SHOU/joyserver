package internal

import (
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
)

type ClientHandler struct {
	Node *Node
}

func (ch *ClientHandler) Register(client *net.Client) {
	client.OnNewAgent = ch.NewAgent
	client.OnCloseAgent = ch.CloseAgent

}

func (ch *ClientHandler) NewAgent(agent *net.ClientAgent) {
	ch.Node.AgentManager.AddAgent(agent.RemoteAddr().String(), agent)

	ch.RegisterToCenter(agent)
}

func (ch *ClientHandler) CloseAgent(agent *net.ClientAgent) {
	ch.Node.AgentManager.RemoveAgent(agent.RemoteAddr().String())
}

func (ch *ClientHandler) RegisterToCenter(agent *net.ClientAgent) {
	agent.WriteMsg(&pb.NodeRegisterReq{NodeType: ch.Node.NodeType, NodeId: uint32(ch.Node.NodeID)})
}
