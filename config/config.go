package config

import (
	"bufio"
	"encoding/json"
	"os"
)

type subscriptionType uint

// Config stores application configurations
type Config struct {
	Link       string `json:"Link"`
	Type       subscriptionType
	ConfigPath string `json:"configPath"`
}

var (
	//Conf global variable that used to store global config
	Conf               Config
	lastLoadedFilePath string
)

// NewConfig create config struct from a file path
func NewConfig(file *os.File) *Config {
	if file != nil {
		scner := bufio.NewScanner(file)
		var bs []byte
		for scner.Scan() {
			bs = append(bs, scner.Bytes()...)
		}

		var conf Config
		err := json.Unmarshal(bs, &conf)
		if err != nil {
			return nil
		}

		return &conf
	}
	return nil
}

// NewConfigByPath create config by file path
func NewConfigByPath(path string) *Config {
	lastLoadedFilePath = path
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	return NewConfig(file)
}

// ReloadConfig will reload config from LastLoadedFilePath
// return value : true means reloading is successful
func ReloadConfig() bool {
	c := NewConfigByPath(lastLoadedFilePath)
	if c != nil {
		Conf = *c
	}
	return c != nil
}

// LoadConfig will load config from path
// return value : true means reloading is successful
func LoadConfig(path string) {
	Conf = *NewConfigByPath(path)
}
