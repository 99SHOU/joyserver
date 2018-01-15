package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_struct"
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
	server.Register(pb.EGameMsgID_EGMI_LOGIN_RESPON, &pb.LoginRespon{})
	server.RegisterAndSetHandler(pb.EGameMsgID_EGMI_LOGIN_REQ, &pb.LoginReq{}, handler.onLoginReq)
}

func (handler *ServerHandler) NewAgent(agent *base.Agent) {
}

func (handler *ServerHandler) CloseAgent(agent *base.Agent) {
}

func (handler *ServerHandler) onLoginReq(message interface{}, agent interface{}) {
	msg := message.(*pb.LoginReq)
	a := agent.(*base.Agent)

	token := ""
	gateAddr := ""
	responCode := pb.LoginResponCode_LOGIN_FAIL
	loginRespon := &pb.LoginRespon{Account: msg.Account, Token: token, GateAddr: gateAddr, ResponCode: responCode}

	req := &rpc_struct.VerifyAccountReq{Account: msg.Account}
	resp := &rpc_struct.VerifyAccountRespon{}
	handler.module.RpcMgr.RpcCallByServerType(pb.SERVER_TYPE_CENTER, "RpcHandler.VerifyAccount", &req, &resp)
	if resp.ResponCode == define.VERIFY_ACCOUNT_RESPOND_CODE_SUCCESS {
		req := &rpc_struct.GateInfoReq{Account: msg.Account}
		resp := &rpc_struct.GateInfoRespon{}
		handler.module.RpcMgr.RpcCallByServerType(pb.SERVER_TYPE_CENTER, "RpcHandler.GetGateServerInfo", &req, &resp)
		if resp.ResponCode == define.GET_GATE_ADDR_RESPOND_CODE_SUCCESS {
			responCode = pb.LoginResponCode_LOGIN_SUCCESS

			loginRespon.Account = msg.Account
			loginRespon.GateAddr = resp.GateAddr
			loginRespon.Token = resp.Token
			loginRespon.ResponCode = responCode
		}
	}

	a.WriteMsg(loginRespon)
	return
}
