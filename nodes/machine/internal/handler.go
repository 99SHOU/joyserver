package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/rpc_struct"
)

type RpcHandler struct {
	module *Module
}

func (handler *RpcHandler) ModulePortReq(req *rpc_struct.ModulePortReq, resp *rpc_struct.ModulPortResp) error {
	rpcPort, port, err := handler.module.GetModulePort(req.ServerType)
	if err != nil {
		log.Error("%v", err)
		resp.RpcPort = 0
		resp.ServerType = req.ServerType
	}

	log.Debug("RPC_PORT%v, PORT %v, %v", rpcPort, port, req.ServerType.String())

	resp.RpcPort = rpcPort
	resp.Port = port
	resp.ServerType = req.ServerType

	return err
}
