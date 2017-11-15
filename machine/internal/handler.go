package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/rpc_struct"
)

type RpcHandler struct {
	module *Module
}

func (handler *RpcHandler) ModulePortReq(req *rpc_struct.ModulePortReq, resp *rpc_struct.ModulPortResp) error {
	port, err := handler.module.GetModulePort(req.ServerType)
	if err != nil {
		log.Error("%v", err)
		resp.Port = 0
		resp.ServerType = req.ServerType
	}

	log.Debug("ModulePortReq%v   %v", port, req.ServerType.String())

	resp.Port = port
	resp.ServerType = req.ServerType

	return err
}
