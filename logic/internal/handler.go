package internal

import (
	"github.com/99SHOU/joyserver/common/rpc_struct"
)

type RpcHandler struct {
	module *Module
}

func (handler *RpcHandler) ModulePortReq(req *rpc_struct.ModulePortReq, resp *rpc_struct.ModulPortResp) error {

	return nil
}
