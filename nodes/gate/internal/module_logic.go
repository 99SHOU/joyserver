package internal

import (
	"math/rand"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_struct"
	"time"
)

func (m *Module) GetRandomString(strLen int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < strLen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func (m *Module) GetBalanceLogicServer() uint32 {
	req := &rpc_struct.BalanceLogicServerReq{}
	resp := &rpc_struct.BalanceLogicServerRespon{}
	m.RpcMgr.RpcCallByServerType(pb.SERVER_TYPE_CENTER, "RpcHandler.GetBalanceLogicServer", req, resp)

	if resp.ResponCode == define.COMMON_RESPOND_CODE_SUCCESS {
		return resp.ModuleId
	}

	return 0
}

func (m *Module) AccountEnter(account string) bool {
	req := &rpc_struct.AccountEnter{Account: account}
	resp := &rpc_struct.CommonRespon{}
	m.RpcMgr.RpcCallByServerType(pb.SERVER_TYPE_CENTER, "RpcHandler.AccountEnter", req, resp)

	if resp.ResponCode == define.COMMON_RESPOND_CODE_SUCCESS {
		return true
	}

	return false
}
