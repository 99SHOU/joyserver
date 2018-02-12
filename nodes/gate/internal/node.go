package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/modules"
)

type Node struct {
	base.Node
	ModuleManager modules.ModulesManager

	nodeClientManager *modules.NodeClientManager
}

func (n *Node) OnInit() {
	n.NodeType = pb.NODE_TYPE_GATE
	n.NodeStatu = pb.NODE_STATU_NOT_READY
	n.NodeID = n.NodeCfg.NodeID

	// node模块注册
	n.ModuleManager = modules.NewModulesManager()
	n.ModuleManager.Register(&modules.AgentManager{})
	n.ModuleManager.Register(&modules.NodeClientManager{})
	n.ModuleManager.Register(&TokenMgr{})

	n.ModuleManager.Init()
	n.ModuleManager.AfterInit()

	n.nodeClientManager = n.ModuleManager.Find("NodeClientManager").(*modules.NodeClientManager)

	// 启动服务
	n.NodeServer = net.NewServer(n.NodeCfg.NodePort, &NodeServerHandler{Node: n}, net.NewProcessor())
	n.NodeServer.Start()
}

func (n *Node) OnDestroy() {
	// 关闭服务
	n.NodeServer.Close()

	// node模块销毁
	n.ModuleManager.BeforeDestroy()
	n.ModuleManager.Destroy()
}

func (n *Node) Run(closeSig chan bool) {
	// 连接到center服务器
	n.nodeClientManager.NewAndStart(n.NodeCfg.CenterAddr, &NodeClientHandler{Node: n}, net.NewProcessor())

	for {
		n.ModuleManager.Run()

		if <-closeSig {
			break
		}
	}
}
