package base

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
)

type NodeConfig struct {
	NodeID     define.NodeID
	CenterAddr string
	NodePort   uint32
	ServerPort uint32
}

type Node struct {
	NodeType   pb.NODE_TYPE
	NodeStatu  pb.NODE_STATU
	NodeID     define.NodeID
	NodeCfg    NodeConfig
	NodeServer net.Server
}
