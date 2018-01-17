package internal

import (
	"github.com/99SHOU/joyserver/common/net"
)

type MessageHandler struct {
	Node *Node
}

func (handler *MessageHandler) Register(server net.Server) {
	server.OnNewAgent = handler.NewAgent
	server.OnCloseAgent = handler.CloseAgent

}

func (handler *MessageHandler) NewAgent(agent *net.ServerAgent) {
}

func (handler *MessageHandler) CloseAgent(agent *net.ServerAgent) {
}
