package main

import (
	"flag"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/conf"
	"github.com/99SHOU/joyserver/common/define"
	"github.com/99SHOU/joyserver/nodes/center"
	game "github.com/99SHOU/joyserver/nodes/game"
	"github.com/99SHOU/joyserver/nodes/gate"
	"github.com/99SHOU/joyserver/nodes/login"
	"github.com/99SHOU/joyserver/nodes/machine"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	nodeID := flag.Uint("nodeid", 0, "")
	server := flag.String("server", "error", "")
	centerAddr := flag.String("center", "", "")
	port := flag.Uint("port", 0, "")

	nodeCfg := base.NodeConfig{NodeID: define.NodeID(*nodeID), CenterAddr: *centerAddr, NodePort: uint32(*port)}

	if *server == "error" {
		machine.Node.NodeCfg = base.NodeConfig{NodeID: 0, CenterAddr: "127.0.0.1:2001", NodePort: 2000}
		center.Node.NodeCfg = base.NodeConfig{NodeID: 1, CenterAddr: "127.0.0.1", NodePort: 2001}
		game.Node.NodeCfg = base.NodeConfig{NodeID: 2, CenterAddr: "127.0.0.1:2001", NodePort: 2002}
		gate.Node.NodeCfg = base.NodeConfig{NodeID: 3, CenterAddr: "127.0.0.1:2001", NodePort: 2003}
		login.Node.NodeCfg = base.NodeConfig{NodeID: 4, CenterAddr: "127.0.0.1:2001", NodePort: 2004}

		game.Node1.NodeCfg = base.NodeConfig{NodeID: 5, CenterAddr: "127.0.0.1:2001", NodePort: 2005}
		gate.Node1.NodeCfg = base.NodeConfig{NodeID: 6, CenterAddr: "127.0.0.1:2001", NodePort: 2006}
		game.Node2.NodeCfg = base.NodeConfig{NodeID: 7, CenterAddr: "127.0.0.1:2001", NodePort: 2007}
		gate.Node2.NodeCfg = base.NodeConfig{NodeID: 8, CenterAddr: "127.0.0.1:2001", NodePort: 2008}
		game.Node3.NodeCfg = base.NodeConfig{NodeID: 9, CenterAddr: "127.0.0.1:2001", NodePort: 2009}
		gate.Node3.NodeCfg = base.NodeConfig{NodeID: 10, CenterAddr: "127.0.0.1:2001", NodePort: 2010}

		leaf.Run(
			machine.Node,
			center.Node,
			game.Node,
			gate.Node,
			login.Node,

			// game.Node1,
			// gate.Node1,
			// game.Node2,
			// gate.Node2,
			// game.Node3,
			// gate.Node3,
		)
	} else {
		switch *server {
		case "machine":
			machine.Node.NodeCfg = nodeCfg
			leaf.Run(machine.Node)
		case "center":
			center.Node.NodeCfg = nodeCfg
			leaf.Run(center.Node)
		case "gate":
			gate.Node.NodeCfg = nodeCfg
			leaf.Run(gate.Node)
		case "game":
			game.Node.NodeCfg = nodeCfg
			leaf.Run(game.Node)
		case "login":
			login.Node.NodeCfg = nodeCfg
			leaf.Run(login.Node)
		default:
			log.Error("error server type")
		}
	}
}
