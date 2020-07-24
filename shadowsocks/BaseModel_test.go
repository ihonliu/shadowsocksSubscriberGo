package shadowsocksModels

import (
	"log"
	"net/url"
	"testing"
)

func compare(s0, s1 *SSArgs) [][3]string {
	var result [][3]string = make([][3]string, 0)
	if s0.Server != s1.Server {
		result = append(result, [3]string{"server"})
	}
	if s0.Port != s1.Port {
		result = append(result, [3]string{"Port", s0.Port, s1.Port})
	}
	if s0.Password != s1.Password {
		result = append(result, [3]string{"Password", s0.Password, s1.Password})
	}
	if s0.Method != s1.Method {
		result = append(result, [3]string{"Method", s0.Method, s1.Method})
	}
	if s0.Plugin != s1.Plugin {
		result = append(result, [3]string{"Plugin", s0.Plugin, s1.Plugin})
	}
	if s0.Remarks != s1.Remarks {
		result = append(result, [3]string{"Remarks", s0.Remarks, s1.Remarks})
	}
	if s0.Group != s1.Group {
		result = append(result, [3]string{"Group", s0.Group, s1.Group})
	}

	return result
}

func TestSSArgs(t *testing.T) {
	sStandard := &SSArgs{
		Server:   "test.com",
		Port:     "9988",
		Password: "qAdUkm",
		Method:   "chacha20-ietf-poly1305",
		Plugin:   "obfs-local --obfs=http --obfs-host=49.hk",
		Group:    "ing - 香港 H",
		Remarks:  "ing",
	}
	testCase :=
		`ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTpxQWRVa20@test.com:9988/?plugin=obfs-local%3Bobfs%3Dhttp%3Bobfs-host%3D49.hk&group=aW5n#ing%20-%20%E9%A6%99%E6%B8%AF%20H`
	ul, err := url.ParseRequestURI(testCase)
	if err != nil {
		t.Error("An error occured during parseRequestURI", err)
	}
	ss := &SSArgs{}
	err = ss.AssignFromUrl(ul)
	if err != nil {
		t.Error("An error occured during parseRequestURI", err)
	}
	res := compare(ss, sStandard)
	for _, v := range res {
		t.Errorf("field %s mismatch, got: %s, expected: %s\n", v[0], v[1], v[2])
	}
	// ss
	log.Println(ss)
	ss.Save(".")
}
