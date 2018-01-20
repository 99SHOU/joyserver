package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
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

func (handler *ServerHandler) onLoginReq(message interface{}, agent interface{}) {
	msg := message.(*pb.LoginReq)
	a := agent.(*net.BaseAgent)

	token := ""
	gateAddr := ""
	responCode := pb.LoginResponCode_LOGIN_FAIL
	loginRespon := &pb.LoginRespon{Account: msg.Account, Token: token, GateAddr: gateAddr, ResponCode: responCode}

	a.WriteMsg(loginRespon)
	return
}
