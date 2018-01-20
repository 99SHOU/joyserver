package modules

import (
	"github.com/99SHOU/joyserver/common/base"
)

type NodeInfo struct {
	NodeID   base.NodeID
	NodeType base.NodeType

	NodeAddr string
}

type NodeBroadCaseManager struct {
}
