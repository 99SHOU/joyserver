package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	// "github.com/99SHOU/joyserver/common/pb"
)

type NodeServerHandler struct {
	Node *Node
}

func (h *NodeServerHandler) Register(server *net.Server) {
	server.OnNewAgent = h.NewAgent
	server.OnCloseAgent = h.CloseAgent

	msg.RegisterMsg(server.Processor)
}

func (h *NodeServerHandler) NewAgent(agent net.Agent) {
}

func (h *NodeServerHandler) CloseAgent(agent net.Agent) {
}

func (h *NodeServerHandler) onLoginReq(message interface{}, agent interface{}) {
	// msg := message.(*pb.LoginReq)
	// a := agent.(*net.BaseAgent)

	// token := ""
	// gateAddr := ""
	// responCode := pb.LoginResponCode_LOGIN_FAIL
	// loginRespon := &pb.LoginRespon{Account: msg.Account, Token: token, GateAddr: gateAddr, ResponCode: responCode}

	// a.WriteMsg(loginRespon)
	// return
}
