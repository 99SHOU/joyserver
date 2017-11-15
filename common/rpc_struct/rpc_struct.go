package rpc_struct

import (
	"github.com/99SHOU/joyserver/common/define"
)

//Server To Machine Start
type ModulePortReq struct {
	ServerType define.SERVER_TYPE
}

type ModulPortResp struct {
	ServerType define.SERVER_TYPE
	Port       int
}

//Server To Start End

//Server To Center Start
type ModuleIdReq struct {
	ServerType define.SERVER_TYPE
}

type ModuleIdResp struct {
	ServerType define.SERVER_TYPE
	ModuleId   int
}

type RegisterToCenterReq struct {
	ServerType define.SERVER_TYPE
	ServerAddr string
	ModuleId   int
}

type RegisterToCenterResp struct {
	ModuleId   int
	ResponCode define.REGISTER_RESPOND_CODE
}

//Server To Center End
