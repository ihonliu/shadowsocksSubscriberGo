package shadowsocksModels

import "net/url"

type SSArgsSlice struct {
	members []SSArgs
}

func (ls *SSArgsSlice) Append(link *SSArgs) {
	ls.members = append(ls.members, *link)
}

func (ls *SSArgsSlice) AssignFromURL(urls []*url.URL) {
	for _, v := range urls {
		ls.Append(NewSSArgs(v))
	}
}
