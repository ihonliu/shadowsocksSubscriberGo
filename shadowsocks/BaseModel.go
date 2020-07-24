// Shadowsocks URI data model
// NOTICE:
// this data model is not fully based on SIP002.
// My focus is mainly based on the service the provide
// and this model is compactible with Catchflying(service provider)
package shadowsocksModels

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type serverInfo struct {
	group string
	name  string
}

// parsePluginArgument will parse plugin argument string to struct we needed
// For example, we have "obfs-local;obfs=http;obfs-host=example.com"
// then this function will convert it to
// pluginArgs{name:obfs-local,opts:obfs=http;obfs-host=example.com}
// Here I have something to declare:
// args field will be empty as it is, just because I do not know how to map
// it with shadowsocks argument, so I will leave it alone. Or maybe sometime
// when I understand how to use it, this field will be fulfilled
func parsePluginArguments(ar string) (result string) {
	ss := strings.Split(ar, ";")
	for _, v := range ss {
		if !strings.Contains(v, "=") {
			result += v + " "
		} else {
			result += "--" + v + " "
		}
	}
	result = strings.TrimRight(result, " ")
	return
}

func parseGroupArguments(ar string) (result serverInfo, err error) {
	ss := strings.Split(ar, "#")
	bs, err := base64.RawStdEncoding.DecodeString(ss[0])
	if err != nil {
		return
	}
	result.group = ss[1]
	result.name = string(bs)
	return
}

//SSArgs implements shadowsocks uri model
type SSArgs struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Method   string `json:"method"`
	Plugin   string `json:"plugin"`
	Remarks  string `json:"remarks"`
	Group    string `json:"-"`
}

// AssignFromUrl will copy settings from give url.URL struct
func (l *SSArgs) AssignFromUrl(ul *url.URL) error {
	l.Server = ul.Hostname()
	l.Port = ul.Port()
	bs, err := base64.RawStdEncoding.DecodeString(ul.User.Username())
	if err != nil {
		return err
	}
	ss := strings.Split(string(bs), ":")
	l.Method = ss[0]
	l.Password = ss[1]
	m := ul.Query()
	if len(m["plugin"]) == 1 {
		l.Plugin = parsePluginArguments(m["plugin"][0])
	}
	if len(m["group"]) == 1 {
		pga, err := parseGroupArguments(m["group"][0])
		if err == nil {
			l.Remarks = pga.name
			l.Group = pga.group
		}
	}
	return nil
}

func (l *SSArgs) Save(path string) error {
	bs, err := json.Marshal(*l)
	if err != nil {
		return err
	}
	tmp := strings.ReplaceAll(l.Group, " - ", "_")
	tmp = strings.ReplaceAll(tmp, " ", "_")
	if path == "" {
		path, err = filepath.Abs(os.Args[0])
		if err != nil {
			return errors.New("Save config failed\n" + err.Error())
		}
		path = filepath.Join(filepath.Dir(path), tmp+".json")
	} else {
		path = filepath.Join(path, tmp+".json")
		fmt.Println(path)
	}
	err = ioutil.WriteFile(path, bs, 0640)
	if err != nil {
		return err
	}
	// log.Println("file saved")
	return nil
}

func NewSSArgs(u *url.URL) *SSArgs {
	result := &SSArgs{}
	if err := result.AssignFromUrl(u); err != nil {
		return nil
	}
	return result
}
