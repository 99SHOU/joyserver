package internal

import (
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/modules"
)

type Node struct {
	base.Node
	AgentManager  modules.AgentManager
	ClientManager modules.NodeClientManager
}

func (n *Node) OnInit() {
	n.NodeType = pb.NODE_TYPE_LOGIN
	n.NodeStatu = pb.NODE_STATU_NOT_READY
	n.NodeID = n.NodeCfg.NodeID

	// node模块初始化
	n.AgentManager = modules.AgentManager{}
	n.ClientManager = modules.NodeClientManager{}
	n.AgentManager.Init()
	n.ClientManager.Init()

	// 启动服务
	n.NodeServer = net.NewServer(n.NodeCfg.NodePort, &NodeServerHandler{Node: n}, net.NewProcessor())
	n.NodeServer.Start()
}

func (n *Node) OnDestroy() {
	// 关闭服务
	n.NodeServer.Close()

	// node模块销毁
	n.ClientManager.Destroy()
	n.AgentManager.Destroy()
}

func (n *Node) Run(closeSig chan bool) {
	// 连接到center服务器
	n.ClientManager.NewAndStart(n.NodeCfg.CenterAddr, &NodeClientHandler{Node: n}, net.NewProcessor())

	for {
		n.AgentManager.Run()
		n.ClientManager.Run()

		if <-closeSig {
			break
		}
	}
}
