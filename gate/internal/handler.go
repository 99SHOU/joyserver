package internal

import (
	"server/common/rpc_struct"
)

type RpcHandler struct {
	module *Module
}

func (handler *RpcHandler) ModulePortReq(req *rpc_struct.ModulePortReq, resp *rpc_struct.ModulPortResp) error {

	return nil
}
