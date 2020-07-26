package functions

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/ihonliu/shadowsocksSubscriberGo/config"
	shadowsocksModels "github.com/ihonliu/shadowsocksSubscriberGo/shadowsocks"
)

// Download will download file from given url
func Download(url string) error {
	if !validURL(url) {
		return ErrURLNotValid
	}
	ctx := context.Background()
	// ctx, cancFunc := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	// defer cancFunc()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	var response []byte
	return httpDo(ctx, req, func(r *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer r.Body.Close()
		response, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		if config.Conf.SaveResponse {
			saveResponse(response)
		}
		ul := decryptSuscription(response)
		tmpSSArgs := &shadowsocksModels.SSArgs{}
		// absPath, err := filepath.Abs(os.Args[0])
		path, err := filepath.Abs(config.Conf.SavePath)
		if err != nil {
			return err
		}
		// path := filepath.Join(filepath.Dir(absPath), config.Conf.SavePath)
		fmt.Println(path)
		os.MkdirAll(path, os.ModePerm)
		for _, v := range ul {
			tmpSSArgs.AssignFromUrl(v)
			tmpSSArgs.Save(path)
		}
		return nil
	})
}

func saveResponse(resp []byte) {
	err := ioutil.WriteFile("result.sss", resp, 0640)
	if err != nil {
		log.Panic("An error occured during saving response\n", err)
	}
	log.Println("response saved :", "result.sss")
}

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	c := make(chan error, 1)
	req.WithContext(ctx)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36")

	go func() {
		c <- f(http.DefaultClient.Do(req))
	}()
	select {
	case <-ctx.Done():
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}

func validURL(testURL string) bool {
	//ToDo: NEED TO IMPLEMENTED
	_, err := url.ParseRequestURI(testURL)
	return err == nil
}
