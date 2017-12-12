package rpc_struct

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/pb"
)

//Server To Machine Start
type ModulePortReq struct {
	ServerType pb.SERVER_TYPE
}

type ModulPortResp struct {
	ServerType pb.SERVER_TYPE
	RpcPort    uint32
	Port       uint32
}

//Server To Start End

//Server To Center Start
type ModuleIdReq struct {
	ServerType pb.SERVER_TYPE
}

type ModuleIdResp struct {
	ModuleId   uint32
	ServerType pb.SERVER_TYPE
}

type RegisterToCenterReq struct {
	ModuleId   uint32
	ServerType pb.SERVER_TYPE
	Ip         string
	RpcPort    uint32
	Port       uint32
}

type RegisterToCenterResp struct {
	ModuleId   uint32
	ServerType pb.SERVER_TYPE
	ResponCode define.REGISTER_RESPOND_CODE
}

// //Server To Center End

type CommonRespon struct {
	ResponCode define.COMMON_RESPOND_CODE
}

type VerifyAccountReq struct {
	Account string
}

type VerifyAccountRespon struct {
	Account    string
	ResponCode define.VERIFY_ACCOUNT_RESPOND_CODE
}

type GateInfoReq struct {
	Account string
}

type GateInfoRespon struct {
	Account    string
	GateAddr   string
	Token      string
	ResponCode define.GET_GATE_ADDR_RESPOND_CODE
}

type GateTokenReq struct {
	Account string
}

type GateTokenRespon struct {
	Account    string
	Token      string
	ResponCode define.GATE_TOKEN_RESPOND_CODE
}

type BalanceLogicServerReq struct {
}

type BalanceLogicServerRespon struct {
	ModuleId   uint32
	ResponCode define.COMMON_RESPOND_CODE
}


type AccountEnter struct {
	Account string
}
