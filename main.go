package main

import (
	"flag"
	"github.com/name5566/leaf"
	"github.com/name5566/leaf/log"
	lconf "github.com/name5566/leaf/conf"
	"github.com/99SHOU/joyserver/nodes/center"
	"github.com/99SHOU/joyserver/common/conf"
	"github.com/99SHOU/joyserver/nodes/gate"
	"github.com/99SHOU/joyserver/nodes/logic"
	"github.com/99SHOU/joyserver/nodes/login"
	"github.com/99SHOU/joyserver/nodes/machine"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	server := flag.String("server", "error", "")

	if *server == "error" {
		leaf.Run(
			machine.Module,
			center.Module,
			gate.Module,
			logic.Module,
			login.Module,
		)	
	} else {
		switch *server {
		case "machine" :
			leaf.Run(machine.Module)
		case "center" :
			leaf.Run(center.Module)
		case "gate" :
			leaf.Run(gate.Module)
		case "logic" :
			leaf.Run(logic.Module)
		case "login" :
			leaf.Run(login.Module)
		default:
			log.Error("error server type")
		}	
	}
}
