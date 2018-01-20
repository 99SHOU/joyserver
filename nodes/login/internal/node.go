package internal

import (
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
)

type Node struct {
	base.Node
}

func (n *Node) OnInit() {
	n.NodeType = pb.NODE_TYPE_LOGIN
	n.NodeID = n.NodeCfg.NodeID

	n.Server = net.NewServer(n.NodeCfg.Port, &ServerHandler{Node: n}, msg.Processor)
}

func (n *Node) OnDestroy() {
	n.Server.Close()
}

func (n *Node) Run(chan bool) {
	n.Server.Start()

	for {

	}
}