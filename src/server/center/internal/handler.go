package internal

import (
	"github.com/name5566/leaf/log"
	"server/common/define"
	"server/common/rpc_struct"
)

type RpcHandler struct {
	module *Module
}

func (handler *RpcHandler) ModuleIdReq(req *rpc_struct.ModuleIdReq, resp *rpc_struct.ModuleIdResp) error {
	if req.ServerType <= define.SERVER_TYPE_START || req.ServerType >= define.SERVER_TYPE_END {
		resp.ModuleId = 0
		resp.ServerType = req.ServerType
		log.Error("RegisterModule ERROR: invalue ServerType %v", req)
	}

	moduleId := handler.module.moduleIdMgr.GetModuleId()
	resp.ModuleId = moduleId
	resp.ServerType = req.ServerType

	return nil
}

func (handler *RpcHandler) RegisterToCenter(req *rpc_struct.RegisterToCenterReq, resp *rpc_struct.RegisterToCenterResp) error {
	if req.ServerType <= define.SERVER_TYPE_START || req.ServerType >= define.SERVER_TYPE_END {
		resp.ModuleId = 0
		resp.ResponCode = define.REGISTER_RESPOND_CODE_FAIL
		log.Error("RegisterModule ERROR: invalue ServerType %v", req)
	}

	moduleClient := handler.module.RpcMgr.NewModuleClient(req.ServerAddr, req.ModuleId, req.ServerType)
	if moduleClient != nil {
		err := handler.module.RpcMgr.AddMoudleClient(moduleClient.ModuleId, moduleClient)
		if err == nil {
			resp.ModuleId = req.ModuleId
			resp.ResponCode = define.REGISTER_RESPOND_CODE_SUCCESS
			return nil
		} else {
			log.Error("%v", err)
		}
	} else {
		log.Error("Connect to machine error")
	}

	resp.ModuleId = req.ModuleId
	resp.ResponCode = define.REGISTER_RESPOND_CODE_FAIL

	return nil
}
