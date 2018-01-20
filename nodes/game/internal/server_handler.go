package internal

import (
	"github.com/99SHOU/joyserver/common/net"
)

type ServerHandler struct {
	Node *Node
}

func (handler *ServerHandler) Register(server net.Server) {
	server.OnNewAgent = handler.NewAgent
	server.OnCloseAgent = handler.CloseAgent

}

func (handler *ServerHandler) NewAgent(agent *net.ServerAgent) {
}

func (handler *ServerHandler) CloseAgent(agent *net.ServerAgent) {
}
