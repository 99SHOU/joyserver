package internal

import (
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
)

var (
	processor = net.NewProcessor()
)

type Node struct {
	base.Node
}

func (n *Node) OnInit() {
	n.NodeType = pb.SERVER_TYPE_MACHINE
	n.NodeID = n.NodeCfg.NodeID

}

func (n *Node) OnDestroy() {

}

func (n *Node) Run(chan bool) {

}
