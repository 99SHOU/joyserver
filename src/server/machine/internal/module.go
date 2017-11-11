package internal

import (
	"server/base"
	"server/common/define"
	"server/common/mgr"
	"server/common/module_client"
	"strconv"
)

type Module struct {
	base.Module
	portMgr *mgr.PortMgr
}

func (m *Module) OnInit() {
	m.ModuleId = define.MACHINE_MODULE_ID
	m.ServerType = define.SERVER_TYPE_MACHINE
	m.portMgr = new(mgr.PortMgr)
	m.RpcMgr = &mgr.RpcMgr{ModuleClient: make(map[int]*module_client.ModuleClient), ServerType: m.ServerType}
	m.RpcHandler = &RpcHandler{module: m}

	m.portMgr.OnInit()
	m.RpcMgr.StartRpcServer(m.RpcHandler, "127.0.0.1:"+strconv.Itoa(define.MACHINE_PORT), m.ModuleId)
}

func (m *Module) OnDestroy() {

}

func (m *Module) Run(chan bool) {

}

func (m *Module) GetModulePort(serverType define.SERVER_TYPE) (int, error) {
	return m.portMgr.GetModulePort(serverType)
}
