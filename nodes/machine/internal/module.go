package internal

import (
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/mgr"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_client"
	"strconv"
)

type Module struct {
	base.Node
	portMgr *mgr.PortMgr
}

func (m *Module) OnInit() {
	m.NodeId = define.MACHINE_MODULE_ID
	m.NodeType = pb.SERVER_TYPE_MACHINE
	m.portMgr = new(mgr.PortMgr)
	m.RpcMgr = &mgr.RpcMgr{RpcClient: make(map[uint32]*rpc_client.RpcClient), ServerType: m.NodeType}
	m.RpcHandler = &RpcHandler{module: m}

	m.portMgr.OnInit()
	m.RpcMgr.StartRpcServer(m.RpcHandler, "127.0.0.1:"+strconv.Itoa(define.MACHINE_RPC_PORT), m.NodeId)
}

func (m *Module) OnDestroy() {

}

func (m *Module) Run(chan bool) {

}

func (m *Module) GetModulePort(serverType pb.SERVER_TYPE) (uint32, uint32, error) {
	return m.portMgr.GetModulePort(serverType)
}
