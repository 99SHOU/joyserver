package module_client

import (
	"fmt"
	"net/rpc"
	"server/common/define"
)

type ModuleClient struct {
	RpcClient  *rpc.Client
	ModuleId   int
	ServerAddr string
	ServerType define.SERVER_TYPE
}

func (mc *ModuleClient) String() string {
	return fmt.Sprintf("ModuleClient: {ModuleId: %v, ServerAddr: %v, ServerType: %v}", mc.ModuleId, mc.ServerAddr, mc.ServerType)
}
