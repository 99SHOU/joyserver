package main

import (
	"flag"
	"github.com/99SHOU/joyserver/common/base"
	"github.com/99SHOU/joyserver/common/conf"
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

	nodeCfg := base.NodeConfig{NodeID: base.NodeID(*nodeID), CenterAddr: *centerAddr, Port: *port}

	if *server == "error" {
		machine.Node.NodeCfg = base.NodeConfig{NodeID: 0, CenterAddr: "127.0.0.1", Port: 2000}
		center.Node.NodeCfg = base.NodeConfig{NodeID: 1, CenterAddr: "127.0.0.1", Port: 2001}
		gate.Node.NodeCfg = base.NodeConfig{NodeID: 2, CenterAddr: "127.0.0.1", Port: 2002}
		game.Node.NodeCfg = base.NodeConfig{NodeID: 3, CenterAddr: "127.0.0.1", Port: 2003}
		login.Node.NodeCfg = base.NodeConfig{NodeID: 4, CenterAddr: "127.0.0.1", Port: 2004}

		leaf.Run(
			machine.Node,
			center.Node,
			gate.Node,
			game.Node,
			login.Node,
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
