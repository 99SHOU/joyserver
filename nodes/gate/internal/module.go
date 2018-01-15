package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/mgr"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/common/rpc_client"
	"strconv"
)

type Module struct {
	base.Node
	*base.Server
	tokenMgr *TokenMgr
}

func (m *Module) OnInit() {
	m.NodeType = pb.SERVER_TYPE_GATE
	m.RpcMgr = &mgr.RpcMgr{RpcClient: make(map[uint32]*rpc_client.RpcClient), ServerType: m.NodeType}
	m.RpcHandler = &RpcHandler{module: m}
	m.ServerHandler = &ServerHandler{module: m}
	m.tokenMgr = NewTokenMgr()

	m.CreateRpcClientToMachine()
	m.CreateRpcClientToCenter()
	m.ModuleIdReq()
	m.ModulePortReq()
	m.StartRpcServer(m.RpcHandler)
	m.Server = base.NewServer("127.0.0.1" + ":" + strconv.FormatUint(uint64(m.Port), 10))
	m.ServerHandler.Init(m.Server)
}

func (m *Module) OnDestroy() {
	m.Server.Close()
}

func (m *Module) Run(chan bool) {
	m.Server.Start()
	if m.RegisterToCenter() == true {
		//log.Error("Connect to Center Error")
	}
}
