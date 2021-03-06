package main

import (
	"github.com/pefish/gene-beauty-ether-addr/cmd/gene-beauty-ether-addr/command"
	"github.com/pefish/gene-beauty-ether-addr/version"
	"github.com/pefish/go-commander"
	go_logger "github.com/pefish/go-logger"
)

func main() {
	commanderInstance := commander.NewCommander(version.AppName, version.Version, version.AppName + " 是一个生成漂亮的以太坊地址的工具，祝你玩得开心。作者：pefish")
	commanderInstance.RegisterSubcommand("inspect", "查看地址信息",command.NewInspectCommand())
	commanderInstance.RegisterDefaultSubcommand("生成地址", command.NewDefaultCommand())
	err := commanderInstance.Run()
	if err != nil {
		go_logger.Logger.Error(err)
	}
}
