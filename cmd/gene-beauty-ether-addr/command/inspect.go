package command

import (
	"flag"
	"github.com/pefish/go-coin-eth"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	go_logger "github.com/pefish/go-logger"
)

type InspectCommand struct {

}

func NewInspectCommand() *InspectCommand {
	return &InspectCommand{

	}
}

func (dc *InspectCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	flagSet.String("path", "m/0/0", "path")
	flagSet.String("mnemonic", "mnemonic", "mnemonic")
	flagSet.String("pass", "pefish", "password")
	return nil
}


func (dc *InspectCommand) OnExited(data *commander.StartData) error {
	return nil
}

func (dc *InspectCommand) Start(data *commander.StartData) error {
	path := go_config.ConfigManagerInstance.MustGetString("path")
	mnemonic := go_config.ConfigManagerInstance.MustGetString("mnemonic")
	pass := go_config.ConfigManagerInstance.MustGetString("pass")

	wallet := go_coin_eth.NewWallet()
	seed := wallet.SeedHexByMnemonic(mnemonic, pass)
	result, err := wallet.DeriveFromPath(seed, path)
	if err != nil {
		return err
	}
	go_logger.Logger.InfoF("Address: %s", result.Address)
	go_logger.Logger.InfoF("Private key: %s", result.PrivateKey)

	return nil
}

