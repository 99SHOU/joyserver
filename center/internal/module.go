package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/base"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/mgr"
	"github.com/99SHOU/joyserver/common/module_client"
)

type Module struct {
	base.Module
	moduleIdMgr *mgr.ModuleIdMgr
}

func (m *Module) OnInit() {
	m.ServerStatu = define.SERVER_STATU_INVALUE
	m.ServerType = define.SERVER_TYPE_CENTER
	m.moduleIdMgr = &mgr.ModuleIdMgr{StartId: define.CENTER_MODULE_ID}
	m.ModuleId = define.CENTER_MODULE_ID
	m.Port = define.CENTER_PORT
	m.RpcMgr = &mgr.RpcMgr{ModuleClient: make(map[int]*module_client.ModuleClient), ServerType: m.ServerType}
	m.RpcHandler = &RpcHandler{module: m}

	m.ServerStatu = define.SERVER_STATU_REFUSE_SERVICE

	m.CreateRpcClientToMachine()
	m.StartRpcServer(m.RpcHandler)

	m.ServerStatu = define.SERVER_STATU_START_SERVICE
}

func (m *Module) OnDestroy() {

}

func (m *Module) Run(chan bool) {

}
