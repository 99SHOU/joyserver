package internal

import (
	//"github.com/name5566/leaf/log"
	"server/base"
	"server/common/define"
	"server/common/mgr"
	"server/common/module_client"
)

type Module struct {
	base.Module
}

func (m *Module) OnInit() {
	m.ServerStatu = define.SERVER_STATU_INVALUE
	m.ServerType = define.SERVER_TYPE_GATE
	m.RpcMgr = &mgr.RpcMgr{ModuleClient: make(map[int]*module_client.ModuleClient), ServerType: m.ServerType}
	m.RpcHandler = &RpcHandler{module: m}

	m.ServerStatu = define.SERVER_STATU_REFUSE_SERVICE

	m.CreateRpcClientToMachine()
	m.CreateRpcClientToCenter()
	m.ModuleIdReq()
	m.ModulePortReq()
	m.StartRpcServer(m.RpcHandler)
	m.RegisterToCenter()

	m.ServerStatu = define.SERVER_STATU_START_SERVICE
}

func (m *Module) OnDestroy() {

}

func (m *Module) Run(chan bool) {

}
