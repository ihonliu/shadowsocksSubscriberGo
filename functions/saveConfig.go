package functions

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ihonliu/shadowsocksSubscriberGo/config"
)

const configExt = ".json"

// SaveConfig will save config stores in config which is passed during initialziation
func SaveConfig(path string) {
	if path == "" {
		var err error
		path, err = filepath.Abs(os.Args[0])
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(path)
		ext := filepath.Ext(path)
		if ext != configExt {
			path = path[0:len(path)-len(ext)] + configExt
		}
	}
	config.Conf.Save(path)
}
