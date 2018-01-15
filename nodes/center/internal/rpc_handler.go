package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_client"
	"github.com/99SHOU/joyserver/common/rpc_struct"
	"strconv"
)

type RpcHandler struct {
	module *Module
}

func (handler *RpcHandler) ModuleIdReq(req *rpc_struct.ModuleIdReq, resp *rpc_struct.ModuleIdResp) error {
	if req.ServerType <= pb.SERVER_TYPE_START || req.ServerType >= pb.SERVER_TYPE_END {
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
	if req.ServerType <= pb.SERVER_TYPE_START || req.ServerType >= pb.SERVER_TYPE_END {
		resp.ModuleId = 0
		resp.ResponCode = define.REGISTER_RESPOND_CODE_FAIL
		log.Error("RegisterModule ERROR: invalue ServerType %v", req)
	}

	rpcClient := handler.module.RpcMgr.NewRpcClient(req.Ip+":"+strconv.FormatUint(uint64(req.RpcPort), 10), req.ModuleId, req.ServerType)
	if rpcClient != nil {
		switch req.ServerType {
		case pb.SERVER_TYPE_LOGIN:
			clientConnectAddr := req.Ip + ":" + strconv.FormatUint(uint64(req.Port), 10)
			userData := &rpc_client.LoginServerUserData{ConnectCount: 0, ClientConnectAddr: clientConnectAddr}
			rpcClient.SetUserData(*userData)
		case pb.SERVER_TYPE_GATE:
			clientConnectAddr := req.Ip + ":" + strconv.FormatUint(uint64(req.Port), 10)
			userData := &rpc_client.GateServerUserData{ConnectCount: 0, ClientConnectAddr: clientConnectAddr}
			rpcClient.SetUserData(*userData)
		}

		err := handler.module.RpcMgr.AddMoudleClient(rpcClient.ModuleId, rpcClient)
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

func (handler *RpcHandler) VerifyAccount(req *rpc_struct.VerifyAccountReq, resp *rpc_struct.VerifyAccountRespon) error {
	ok, err := handler.module.accountVerifyMgr.VarifyAccount(req.Account)
	if err != nil {
		log.Error(err.Error())
	}

	if ok {
		resp.Account = req.Account
		resp.ResponCode = define.VERIFY_ACCOUNT_RESPOND_CODE_SUCCESS
		return nil
	}

	resp.Account = req.Account
	resp.ResponCode = define.VERIFY_ACCOUNT_RESPOND_CODE_FAIL

	return nil
}

func (handler *RpcHandler) GetGateServerInfo(req *rpc_struct.GateInfoReq, resp *rpc_struct.GateInfoRespon) error {
	gateConnectAddr, token, err := handler.module.GetBalanceGateAddr(req.Account)
	if err == nil {
		resp.Account = req.Account
		resp.GateAddr = gateConnectAddr
		resp.Token = token
		resp.ResponCode = define.GET_GATE_ADDR_RESPOND_CODE_SUCCESS
	} else {
		resp.Account = req.Account
		resp.GateAddr = ""
		resp.Token = ""
		resp.ResponCode = define.GET_GATE_ADDR_RESPOND_CODE_FAIL
		log.Error(err.Error())
	}

	return nil
}

// func (handler *RpcHandler) GetLogicServerInfo(rep *rpc_struct.BalanceLogicServerReq, resp *rpc_struct.BalanceLogicServerRespon) error {

// }
