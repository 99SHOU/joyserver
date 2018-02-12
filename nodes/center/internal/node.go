package internal

import (
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/modules"
)

type Node struct {
	base.Node
	ModulesManager modules.ModulesManager
}

func (n *Node) OnInit() {
	n.NodeType = pb.NODE_TYPE_CENTER
	n.NodeStatu = pb.NODE_STATU_NOT_READY
	n.NodeID = n.NodeCfg.NodeID

	// node模块注册
	n.ModulesManager = modules.NewModulesManager()
	n.ModulesManager.Register(&modules.DBManager{})
	n.ModulesManager.Register(&modules.AgentManager{})

	n.ModulesManager.Init()
	n.ModulesManager.AfterInit()

	n.NodeStatu = pb.NODE_STATU_READY

	// 启动服务
	n.NodeServer = net.NewServer(n.NodeCfg.NodePort, &NodeServerHandler{Node: n}, net.NewProcessor())
	n.NodeServer.Start()
}

func (n *Node) OnDestroy() {
	// 关闭服务
	n.NodeServer.Close()

	// node模块销毁
	n.ModulesManager.BeforeDestroy()
	n.ModulesManager.Destroy()
}

func (n *Node) Run(closeSig chan bool) {

	for {
		n.ModulesManager.Run()

		if <-closeSig {
			break
		}
	}
}
