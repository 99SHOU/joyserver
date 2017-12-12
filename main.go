package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"github.com/99SHOU/joyserver/center"
	"github.com/99SHOU/joyserver/common/conf"
	"github.com/99SHOU/joyserver/gate"
	"github.com/99SHOU/joyserver/logic"
	"github.com/99SHOU/joyserver/login"
	"github.com/99SHOU/joyserver/machine"
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
