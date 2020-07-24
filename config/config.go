package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type subscriptionType uint

// Config stores application configurations
type Config struct {
	Link         string `json:"Link"`
	Type         subscriptionType
	ConfigPath   string `json:"-"`
	SaveResponse bool   `json:"saveResponse"`
	SavePath     string `json:"savePath"`
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

func (c *Config) Save(path string) error {
	bs, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bs, 0640)
	if err != nil {
		return err
	}
	log.Println("Config saved successfully")
	return nil
}

func (c *Config) Load(path string) error {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, &Conf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Check() error {
	if c.Link == "" {
		return errors.New("Link field cannot be empty")
	}
	return nil

}
