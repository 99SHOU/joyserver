package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	// "github.com/99SHOU/joyserver/common/pb"
)

type ServerHandler struct {
	Node *Node
}

func (sh *ServerHandler) Register(server *net.Server) {
	server.OnNewAgent = sh.NewAgent
	server.OnCloseAgent = sh.CloseAgent

	msg.RegisterMsg(server.Processor)
}

func (sh *ServerHandler) NewAgent(agent net.Agent) {
}

func (sh *ServerHandler) CloseAgent(agent net.Agent) {
}

func (sh *ServerHandler) onLoginReq(message interface{}, agent interface{}) {
	// msg := message.(*pb.LoginReq)
	// a := agent.(*net.BaseAgent)

	// token := ""
	// gateAddr := ""
	// responCode := pb.LoginResponCode_LOGIN_FAIL
	// loginRespon := &pb.LoginRespon{Account: msg.Account, Token: token, GateAddr: gateAddr, ResponCode: responCode}

	// a.WriteMsg(loginRespon)
	// return
}
