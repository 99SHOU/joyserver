package internal

import (
	"github.com/99SHOU/joyserver/common/net"
)

type ServerHandler struct {
	Node *Node
}

func (sh *ServerHandler) Register(server *net.Server) {
	server.OnNewAgent = sh.NewAgent
	server.OnCloseAgent = sh.CloseAgent

}

func (sh *ServerHandler) NewAgent(agent net.Agent) {
}

func (sh *ServerHandler) CloseAgent(agent net.Agent) {
}
