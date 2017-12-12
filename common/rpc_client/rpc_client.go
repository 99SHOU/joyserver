package rpc_client

import (
	"fmt"
	"net/rpc"
	"github.com/99SHOU/joyserver/common/pb"
)

type LoginServerUserData struct {
	ClientConnectAddr string
	ConnectCount      int64
}

type GateServerUserData struct {
	ClientConnectAddr string
	ConnectCount      int64
}

type MachineServerUserData struct {
}

type CenterServerUserData struct {
}

type LogicServerUserData struct {
	ConnectCount int64
}

type RpcClient struct {
	RpcClient  *rpc.Client
	ModuleId   uint32
	ServerAddr string
	ServerType pb.SERVER_TYPE
	userData   interface{}
}

func (mc *RpcClient) String() string {
	return fmt.Sprintf("RpcClient: {ModuleId: %v, ServerAddr: %v, ServerType: %v}", mc.ModuleId, mc.ServerAddr, mc.ServerType)
}

func (mc *RpcClient) UserData() interface{} {
	return mc.userData
}

func (mc *RpcClient) SetUserData(userData interface{}) {
	mc.userData = userData
}
