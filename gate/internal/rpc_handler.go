package internal

import (
	"fmt"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/rpc_struct"
	"strconv"
	"time"
)

type RpcHandler struct {
	module *Module
}

func (handler *RpcHandler) ModulePortReq(req *rpc_struct.ModulePortReq, resp *rpc_struct.ModulPortResp) error {

	return nil
}

func (handler *RpcHandler) GateTokenReq(req *rpc_struct.GateTokenReq, resp *rpc_struct.GateTokenRespon) error {
	randomString := handler.module.GetRandomString(define.GATE_RANDOM_STRING_LEN)
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	token := fmt.Sprintf("%v#%v#%v", req.Account, randomString, timeStamp)

	handler.module.tokenMgr.AddTokenInfo(req.Account, token)

	resp.Account = req.Account
	resp.Token = token
	resp.ResponCode = define.GATE_TOKEN_RESPOND_CODE_SUCCESS

	return nil
}
