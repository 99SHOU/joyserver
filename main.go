package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/center"
	"server/conf"
	"server/gate"
	"server/logic"
	"server/login"
	"server/machine"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		machine.Module,
		center.Module,
		gate.Module,
		logic.Module,
		login.Module,
	)
}
