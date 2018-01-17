package base

import (
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
)

type NodeID uint32

type NodeConfig struct {
	NodeID     NodeID
	CenterAddr string
	Port       uint
}

type Node struct {
	NodeType pb.SERVER_TYPE
	NodeID   NodeID
	NodeCfg  NodeConfig
	Server   net.Server
}
