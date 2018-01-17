package internal

import (
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
}

func (n *Node) OnInit() {
	n.NodeType = pb.SERVER_TYPE_CENTER
	n.NodeID = n.NodeCfg.NodeID

	db := mysql.Open(define.MYSQL_DNS)
	n.accountVerifyMgr = &AccountVerifyMgr{db: db}
	n.accountVerifyMgr.Init()

	n.Server = net.NewServer(n.NodeCfg.Port, &MessageHandler{Node: n}, msg.Processor)
}

func (n *Node) OnDestroy() {
	n.accountVerifyMgr.Destroy()
}

func (n *Node) Run(chan bool) {
	n.Server.Start()
}
