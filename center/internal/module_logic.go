package internal

import (
	"errors"
	"math"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_client"
	"github.com/99SHOU/joyserver/common/rpc_struct"
)

func (m *Module) GetBalanceGateAddr(account string) (string, string, error) {
	moduleId := m.GetBalanceServer(pb.SERVER_TYPE_GATE)
	if moduleId == 0 {
		return "", "", errors.New("Can not find a Gate")
	}

	userData := m.GetUserInfo(moduleId).(rpc_client.GateServerUserData)
	token, err := m.GetGateToken(moduleId, account)

	return userData.ClientConnectAddr, token, err
}

// func (m *Module) GetBalanceLogicAddr() (string, error) {
// 	moduleId := m.GetBalanceServer(pb.SERVER_TYPE_LOGIC)
// 	if moduleId == 0 {
// 		return "", errors.New("Can not find a Logic")
// 	}

// 	userData := m.GetUserInfo(moduleId).(rpc_client.LogicServerUserData)

// 	return userData.
// }

func (m *Module) GetGateToken(moduleId uint32, account string) (string, error) {
	req := &rpc_struct.GateTokenReq{Account: account}
	resp := &rpc_struct.GateTokenRespon{}
	m.RpcMgr.RpcCallByModuleId(moduleId, "RpcHandler.GateTokenReq", req, resp)

	if resp.ResponCode == define.GATE_TOKEN_RESPOND_CODE_SUCCESS {
		return resp.Token, nil
	}

	return "", errors.New("Get gate token error")
}

// maybe move to module_base.go start

func (m *Module) GetBalanceServer(serverType pb.SERVER_TYPE) uint32 {
	var minConnectCount int64 = math.MaxInt64
	//var minConnectLogicAddr string = ""
	var moduleId uint32 = 0
	servers := m.RpcMgr.GetRpcClientByServerType(serverType)
	for _, v := range servers {
		gateUserData := v.UserData().(rpc_client.GateServerUserData)
		if gateUserData.ConnectCount < minConnectCount {
			minConnectCount = gateUserData.ConnectCount
			//minConnectLogicAddr = gateUserData.ClientConnectAddr
			moduleId = v.ModuleId
		}
	}

	return moduleId
}

func (m *Module) GetUserInfo(moduleId uint32) interface{} {
	rpcClient := m.RpcMgr.GetRpcClientByModuleId(moduleId)
	if rpcClient != nil {
		return rpcClient.UserData()
	}

	return nil
}

//maybe move to module_base.go end
