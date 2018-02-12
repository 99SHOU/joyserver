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
	// node 自身信息初始化
	n.NodeType = pb.NODE_TYPE_GAME
	n.NodeStatu = pb.NODE_STATU_NOT_READY
	n.NodeID = n.NodeCfg.NodeID

	// node模块注册
	n.ModulesManager = modules.NewModulesManager()
	n.ModulesManager.Register(&modules.DBManager{})
	// n.ModulesManager.Register(&modules.LoginManager{})
	n.ModulesManager.Register(&modules.AgentManager{})
	// n.ModulesManager.Register(&modules.PlayerIdManager{})

	n.ModulesManager.Init()
	n.ModulesManager.AfterInit()
}

func (n *Node) OnDestroy() {
	// 关闭服务
	n.NodeServer.Close()

	// node模块销毁
	n.ModulesManager.BeforeDestroy()
	n.ModulesManager.Destroy()
}

func (n *Node) Run(closeSig chan bool) {
	// 启动服务
	n.NodeServer = net.NewServer(n.NodeCfg.ServerPort, &NodeServerHandler{Node: n}, net.NewProcessor())
	n.NodeServer.Start()

	for {
		n.ModulesManager.Run()

		if <-closeSig {
			break
		}
	}
}
