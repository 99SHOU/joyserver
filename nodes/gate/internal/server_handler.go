package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"strconv"
	"strings"
	"time"
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

func (handler *ServerHandler) onConnectToGateReq(message interface{}, agent interface{}) {
	m := message.(*pb.ConnectToGateReq)
	a := agent.(*net.BaseAgent)

	tokenInfo := strings.Split(m.Token, "#")
	timeStamp, err := strconv.ParseInt(tokenInfo[2], 10, 64)
	responCode := pb.ConnectToGateResponCode_CONNECT_TO_GATE_FAIL

	if err == nil {
		if time.Now().Unix()-timeStamp < define.GATE_TOKEN_EXPIRY_TIME {
			token := handler.Node.tokenMgr.GetToken(m.Account)
			if strings.Compare(m.Token, token) == 0 {
				responCode = pb.ConnectToGateResponCode_CONNECT_TO_GATE_SUCCESS
			}
		}
	}

	respon := new(pb.ConnectToGateRespon)
	respon.Account = m.Account
	respon.ResponCode = responCode
	a.WriteMsg(respon)
}
