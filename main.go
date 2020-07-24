package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	config "github.com/ihonliu/shadowsocksSubscriberGo/config"
	"github.com/ihonliu/shadowsocksSubscriberGo/functions"

	flag "github.com/spf13/pflag"
)

func main() {
	// flag processing
	var bShowHelp bool
	var bSaveCurrentSetting bool
	flag.StringVar(&config.Conf.Link, "link", "", "subscription link")
	flag.StringVar(&config.Conf.ConfigPath, "config", "", "subscription link")
	flag.BoolVar(&bShowHelp, "help", false, "Show help")
	flag.BoolVar(&bSaveCurrentSetting, "save", false, "Save current setting to file")
	flag.Usage =
		func() {
			fmt.Fprintf(os.Stderr, "Usage of %s: \n", filepath.Base(os.Args[0]))
			flag.PrintDefaults()
			os.Exit(0)
		}
	flag.Parse()
	if bShowHelp {
		flag.Usage()
	}
	// -----------------

	// load config from parameter then read other value
	if config.Conf.ConfigPath != "" {
		config.Conf.Load(config.Conf.ConfigPath)
		// fmt.Println(config.Conf)
	}

	if bSaveCurrentSetting {
		functions.SaveConfig("")
	}

	var err error
	if err = config.Conf.Check(); err != nil {
		log.Println(err)
		flag.Usage()
		os.Exit(1)
	}
	log.Println("Downloading file " + config.Conf.Link)
	err = functions.Download(config.Conf.Link)
	if err != nil {
		log.Println(err)
	}
}
