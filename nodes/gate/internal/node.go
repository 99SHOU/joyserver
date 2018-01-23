package internal

import (
	//"github.com/name5566/leaf/log"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/msg"
	"github.com/99SHOU/joyserver/common/net"
	"github.com/99SHOU/joyserver/common/pb"
	"github.com/99SHOU/joyserver/modules"
)

type Node struct {
	base.Node
	AgentManager  modules.AgentManager
	ClientManager modules.ClientManager
	TokenManager  TokenMgr
}

func (n *Node) OnInit() {
	n.NodeType = pb.NODE_TYPE_GATE
	n.NodeID = n.NodeCfg.NodeID

	// node模块初始化
	n.AgentManager = modules.AgentManager{}
	n.ClientManager = modules.ClientManager{}
	n.TokenManager = TokenMgr{}
	n.AgentManager.Init()
	n.ClientManager.Init()
	n.TokenManager.Init()

	// 启动服务
	n.Server = net.NewServer(n.NodeCfg.Port, &ServerHandler{Node: n}, msg.Processor)
	n.Server.Start()
}

func (n *Node) OnDestroy() {
	// 关闭服务
	n.Server.Close()

	// node模块销毁
	n.TokenManager.Destroy()
	n.ClientManager.Destroy()
	n.AgentManager.Destroy()
}

func (n *Node) Run(closeSig chan bool) {
	// 连接到center服务器
	n.ClientManager.NewAndStart(n.NodeCfg.CenterAddr, &ClientHandler{}, msg.Processor)

	for {
		n.AgentManager.Run()
		n.ClientManager.Run()

		if <-closeSig {
			break
		}
	}
}
