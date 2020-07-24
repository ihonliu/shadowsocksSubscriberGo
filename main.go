package main

import (
	"fmt"
	"os"
	"path/filepath"

	config "github.com/ihonliu/shadowsocksSubscriberGo/config"

	flag "github.com/spf13/pflag"
)

func main() {
	// flag processing
	var bShowHelp bool
	flag.StringVar(&config.Conf.Link, "link", "http://localhost", "subscription link")
	flag.StringVar(&config.Conf.ConfigPath, "config", "http://localhost", "subscription link")
	flag.BoolVar(&bShowHelp, "help", false, "Show help")
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

	fmt.Println(config.Conf.Link)
	// fmt.Println(link)
}
