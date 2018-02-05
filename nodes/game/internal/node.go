package internal

import (
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/modules"
)

type Node struct {
	base.Node
	AgentManager      modules.AgentManager
	NodeClientManager modules.NodeClientManager
}

func (n *Node) OnInit() {
	// node 自身信息初始化
	n.NodeType = pb.NODE_TYPE_GAME
	n.NodeStatu = pb.NODE_STATU_NOT_READY
	n.NodeID = n.NodeCfg.NodeID

	// node模块初始化
	n.AgentManager = modules.AgentManager{}
	n.NodeClientManager = modules.NodeClientManager{}
	n.AgentManager.Init()
	n.NodeClientManager.Init()

	// 启动服务
	n.NodeServer = net.NewServer(n.NodeCfg.NodePort, &NodeServerHandler{Node: n}, net.NewProcessor())
	n.NodeServer.Start()
}

func (n *Node) OnDestroy() {
	// 关闭服务
	n.NodeServer.Close()

	// node模块销毁
	n.NodeClientManager.Destroy()
	n.AgentManager.Destroy()
}

func (n *Node) Run(closeSig chan bool) {
	// 连接到center服务器
	n.NodeClientManager.NewAndStart(n.NodeCfg.CenterAddr, &NodeClientHandler{Node: n}, net.NewProcessor())

	for {
		n.AgentManager.Run()
		n.NodeClientManager.Run()

		if <-closeSig {
			break
		}
	}
}
