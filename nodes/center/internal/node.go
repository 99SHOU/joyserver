package internal

import (
	"database/sql"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/db/mysql"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/modules"
)

type Node struct {
	base.Node
	AgentManager         modules.AgentManager
	AccountVerifyManager AccountVerifyManager
	db                   *sql.DB
}

func (n *Node) OnInit() {
	n.NodeType = pb.NODE_TYPE_CENTER
	n.NodeStatu = pb.NODE_STATU_NOT_READY
	n.NodeID = n.NodeCfg.NodeID

	// node模块初始化
	n.AgentManager = modules.AgentManager{}
	n.AccountVerifyManager = AccountVerifyManager{db: n.db}
	n.AgentManager.Init()
	n.AccountVerifyManager.Init()

	n.db = mysql.Open(define.MYSQL_DNS)

	n.NodeStatu = pb.NODE_STATU_READY

	// 启动服务
	n.NodeServer = net.NewServer(n.NodeCfg.NodePort, &NodeServerHandler{Node: n}, net.NewProcessor())
	n.NodeServer.Start()
}

func (n *Node) OnDestroy() {
	// 关闭服务
	n.NodeServer.Close()

	// node模块销毁
	n.AgentManager.Destroy()
}

func (n *Node) Run(closeSig chan bool) {

	for {
		n.AgentManager.Run()
		n.AccountVerifyManager.Run()

		if <-closeSig {
			break
		}
	}
}
