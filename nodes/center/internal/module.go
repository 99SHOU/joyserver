package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/db/mysql"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/mgr"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_client"
)

type Module struct {
	base.Node
	moduleIdMgr      *mgr.ModuleIdMgr
	accountVerifyMgr *AccountVerifyMgr
}

func (m *Module) OnInit() {
	m.NodeType = pb.SERVER_TYPE_CENTER
	m.moduleIdMgr = &mgr.ModuleIdMgr{StartId: define.CENTER_MODULE_ID}
	m.NodeId = define.CENTER_MODULE_ID
	m.RpcPort = define.CENTER_RPC_PORT
	m.RpcMgr = &mgr.RpcMgr{RpcClient: make(map[uint32]*rpc_client.RpcClient), ServerType: m.NodeType}
	m.RpcHandler = &RpcHandler{module: m}

	db := mysql.Open(define.MYSQL_DNS)
	m.accountVerifyMgr = &AccountVerifyMgr{db: db}
	m.accountVerifyMgr.Init()

	m.CreateRpcClientToMachine()
	m.StartRpcServer(m.RpcHandler)
}

func (m *Module) OnDestroy() {
	m.accountVerifyMgr.Destroy()
}

func (m *Module) Run(chan bool) {

}
