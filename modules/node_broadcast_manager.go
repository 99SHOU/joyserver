package modules

import (
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/common/pb"
)

type NodeInfo struct {
	NodeID   define.NodeID
	NodeType pb.NODE_TYPE

	NodeAddr string
}

type NodeBroadCaseManager struct {
}
