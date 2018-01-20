package internal

import (
	"database/sql"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/db/mysql"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
)

type Node struct {
	base.Node
	accountVerifyMgr *AccountVerifyMgr
	db               *sql.DB
}

func (n *Node) OnInit() {
	n.NodeType = pb.NODE_TYPE_CENTER
	n.NodeID = n.NodeCfg.NodeID

	n.db = mysql.Open(define.MYSQL_DNS)

	n.accountVerifyMgr = &AccountVerifyMgr{db: n.db}
	n.accountVerifyMgr.Init()
	n.Server = net.NewServer(n.NodeCfg.Port, &ServerHandler{Node: n}, msg.Processor)
}

func (n *Node) OnDestroy() {
	n.Server.Close()
	n.accountVerifyMgr.Destroy()

}

func (n *Node) Run(chan bool) {
	n.Server.Start()

	for {

	}
}
