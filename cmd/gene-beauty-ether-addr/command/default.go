package command

import (
	"context"
	"flag"
	"github.com/pefish/go-coin-eth"
	"github.com/pefish/go-commander"
	go_config "github.com/pefish/go-config"
	"github.com/pefish/go-jsvm/pkg/vm"
	go_logger "github.com/pefish/go-logger"
	"github.com/pefish/go-mysql"
	"github.com/pefish/go-random"
	"strconv"
	"strings"
	"sync"
)

type DefaultCommand struct {
	cacheData []*CacheData
}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{
		cacheData: make([]*CacheData, 0),
	}
}

func (dc *DefaultCommand) DecorateFlagSet(flagSet *flag.FlagSet) error {
	flagSet.String("js-file", "rule.js", "js file")
	flagSet.String("pass", "pefish", "password")
	flagSet.String("thread", "3", "thread count")
	return nil
}

type CacheData struct {
	Path     string `json:"path"`
	Index    int    `json:"index"`
	Mnemonic string `json:"mnemonic"`
}

func (dc *DefaultCommand) OnExited(data *commander.StartData) error {
	err := data.Cache.Save(dc.cacheData)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DefaultCommand) Start(data *commander.StartData) error {
	jsFileName := go_config.Config.MustGetString("js-file")
	pass := go_config.Config.MustGetString("pass")
	mysqlConfig := go_config.Config.MustGetMap("mysql")
	threadCount := go_config.Config.MustGetInt("thread")

	go_mysql.MysqlInstance.SetLogger(go_logger.Logger)
	err := go_mysql.MysqlInstance.ConnectWithMap(mysqlConfig)
	if err != nil {
		return err
	}

	_, err = data.Cache.Load(&dc.cacheData)
	if err != nil {
		return err
	}
	//go_logger.Logger.Info(len(dc.cacheData), threadCount)
	cacheLength := len(dc.cacheData)
	if cacheLength < threadCount {
		for i := 0; i < threadCount - cacheLength; i++ {
			mnemonic := go_random.RandomInstance.MustRandomString(16)
			go_logger.Logger.InfoF("生成 mnemonic: %s", mnemonic)
			dc.cacheData = append(dc.cacheData, &CacheData{
				Path:     "m/44'/60'/0'/0/",
				Index:    0,
				Mnemonic: mnemonic,
			})
		}
	}

	wallet := go_coin_eth.NewWallet()
	var wg sync.WaitGroup
	for _, info := range dc.cacheData {
		wg.Add(1)
		go func(info *CacheData) {
			defer wg.Done()
			go_logger.Logger.InfoF("[%s] 开始工作", info.Mnemonic)
			err := dc.findForever(data.ExitCancelCtx, wallet, jsFileName, info, pass)
			if err != nil {
				go_logger.Logger.Error(err)
			}
			go_logger.Logger.InfoF("[%s] 已退出", info.Mnemonic)
		}(info)
	}
	wg.Wait()
	return nil
}

func (dc *DefaultCommand) findForever(ctx context.Context, wallet *go_coin_eth.Wallet, jsFileName string, info *CacheData, pass string) error {
	jsVm, err := vm.NewVmAndLoadWithFile(jsFileName)
	if err != nil {
		return err
	}
	seed := wallet.SeedHexByMnemonic(info.Mnemonic, pass)
	for {
		select {
		case <- ctx.Done():
			return nil
		default:
			path := info.Path + strconv.Itoa(info.Index)
			result, err := wallet.DeriveFromPath(seed, path)
			if err != nil {
				return err
			}
			//go_logger.Logger.Info(strings.ToLower(result.Address))
			jsRunResult, err := jsVm.Run([]interface{}{
				strings.ToLower(result.Address),
			})
			if err != nil || jsRunResult.(bool) == false {
				info.Index += 1
				continue
			}
			// 满足条件
			go_logger.Logger.InfoF("[%s] %s 满足条件", info.Mnemonic, result.Address)
			_, _, err = go_mysql.MysqlInstance.Insert("address", map[string]interface{}{
				"address":  result.Address,
				"path":     `s: = "` + path + `"`,
				"mnemonic": info.Mnemonic,
			})
			if err != nil {
				go_logger.Logger.Error(err)
			}

			info.Index += 1
		}
	}
}
