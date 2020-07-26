package functions

import (
	"testing"

	"github.com/ihonliu/shadowsocksSubscriberGo/config"
)

func TestDownload(t *testing.T) {
	config.Conf.Load("../shadowsocksSubscriberGo.json")
	config.Conf.SaveResponse = true // test save response
	var err error
	err = Download("123")
	if err != ErrURLNotValid {
		t.Error("return value of download from a non-valid url should be EURLNotValid")
	}
	if err = Download("http://ihonliu.xyz"); err != nil {
		t.Error("download function error")
	}
}

func TestValidUrl(t *testing.T) {
	testCase :=
		[]struct {
			uri    string
			expect bool
		}{
			{"http://google.com", true},
			{"googlecom", false},
			{"google.com", false},
			{"", false}}
	for _, v := range testCase {
		if validURL(v.uri) != v.expect {
			t.Errorf("mismatch! uri: %s, expected: %v", v.uri, v.expect)
		}

	}

}
