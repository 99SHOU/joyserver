package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/pb"
	"strconv"
	"strings"
	"time"
)

type ServerHandler struct {
	module *Module
}

func (handler *ServerHandler) Init(server *base.Server) {
	server.OnNewAgent = handler.NewAgent
	server.OnCloseAgent = handler.CloseAgent

	//所有的协议都必须Register无论是监听或者发送
	//Register注册协议
	//SetHandler监听协议
	//RegisterAndSetHandler注册并监听协议
	server.Register(pb.EGameMsgID_EGMI_CONNECT_TO_GATE_RESPON, &pb.ConnectToGateRespon{})
	server.RegisterAndSetHandler(pb.EGameMsgID_EGMI_CONNECT_TO_GATE_REQ, &pb.ConnectToGateReq{}, handler.onConnectToGateReq)
}

func (handler *ServerHandler) NewAgent(agent *base.Agent) {
}

func (handler *ServerHandler) CloseAgent(agent *base.Agent) {
}

func (handler *ServerHandler) onConnectToGateReq(message interface{}, agent interface{}) {
	m := message.(*pb.ConnectToGateReq)
	a := agent.(*base.Agent)

	tokenInfo := strings.Split(m.Token, "#")
	timeStamp, err := strconv.ParseInt(tokenInfo[2], 10, 64)
	responCode := pb.ConnectToGateResponCode_CONNECT_TO_GATE_FAIL

	if err == nil {
		if time.Now().Unix()-timeStamp < define.GATE_TOKEN_EXPIRY_TIME {
			token := handler.module.tokenMgr.GetToken(m.Account)
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
