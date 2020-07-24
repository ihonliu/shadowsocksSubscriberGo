package functions

import (
	"io/ioutil"
	"log"
	"testing"
)

// TestDecrypt this test should use bare eye to identifiy
func TestDecrypt(t *testing.T) {
	bs, err := ioutil.ReadFile("SS_1595586703.txt")
	if err != nil {
		t.Error(err.Error())
	}

	uls := decryptSuscription(bs)
	for _, v := range uls {
		log.Println(v)
	}

}
