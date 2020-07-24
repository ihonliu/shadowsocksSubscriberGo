package functions

import (
	"encoding/base64"
	"log"
	"net/url"
	"strings"
)

func decryptSuscription(bs []byte) []*url.URL {
	buf := make([]byte, len(bs))
	// log.Println(string(bs))
	n, err := base64.StdEncoding.Decode(buf, bs)
	if err != nil {
		log.Panic("An error occured during decrpting reponse\n", err)
	}
	ss := strings.Split(string(buf[:n]), "\n")
	// uris := parseURI(ss)
	// for _, ul := range uris {
	// log.Printf("scheme:%s,userinfo:%s,host:%s,HostName:%s,port:%s,path:%s,query:%v", ul.Scheme, ul.User.Username(), ul.Host, ul.Hostname(), ul.Port(), ul.Path, ul.Query())
	// }
	return parseURI(ss)
}

func parseURI(uris []string) []*url.URL {
	result := make([]*url.URL, 0)
	for _, v := range uris {
		if v == "" {
			continue
		}
		ul, err := url.ParseRequestURI(v)
		if err != nil {
			log.Println(err)
			continue
		}
		result = append(result, ul)
	}
	return result
}
