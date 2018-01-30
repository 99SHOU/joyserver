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
	ClientManager modules.ClientManager
}

func (n *Node) OnInit() {
	// node 自身信息初始化
	n.NodeType = pb.NODE_TYPE_GAME
	n.NodeStatu = pb.NODE_STATU_NOT_READY
	n.NodeID = n.NodeCfg.NodeID

	// node模块初始化
	n.AgentManager = modules.AgentManager{}
	n.ClientManager = modules.ClientManager{}
	n.AgentManager.Init()
	n.ClientManager.Init()

	// 启动服务
	n.Server = net.NewServer(n.NodeCfg.Port, &ServerHandler{Node: n}, net.NewProcessor())
	n.Server.Start()
}

func (n *Node) OnDestroy() {
	// 关闭服务
	n.Server.Close()

	// node模块销毁
	n.ClientManager.Destroy()
	n.AgentManager.Destroy()
}

func (n *Node) Run(closeSig chan bool) {
	// 连接到center服务器
	n.ClientManager.NewAndStart(n.NodeCfg.CenterAddr, &ClientHandler{Node: n}, net.NewProcessor())

	for {
		n.AgentManager.Run()
		n.ClientManager.Run()

		if <-closeSig {
			break
		}
	}
}
