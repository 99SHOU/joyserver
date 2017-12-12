package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/mgr"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_client"
)

type Module struct {
	base.Module
}

func (m *Module) OnInit() {
	m.ServerType = pb.SERVER_TYPE_LOGIC
	m.RpcMgr = &mgr.RpcMgr{RpcClient: make(map[uint32]*rpc_client.RpcClient), ServerType: m.ServerType}
	m.RpcHandler = &RpcHandler{module: m}

	m.CreateRpcClientToMachine()
	m.CreateRpcClientToCenter()
	m.ModuleIdReq()
	m.ModulePortReq()
	m.StartRpcServer(m.RpcHandler)

}

func (m *Module) OnDestroy() {

}

func (m *Module) Run(chan bool) {
	if m.RegisterToCenter() == true {
		//log.Error("Connect to Center Error")
	}
}
