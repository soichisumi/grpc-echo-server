package main

import (
	"bytes"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"github.com/soichisumi/go-util/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	dat, err := ioutil.ReadFile("request.bin")
	logger.Info("", zap.Int64("size", bytes.NewReader(dat).Size()))
	logger.Info("", zap.Int("len", bytes.NewReader(dat).Len()))
	if err != nil {
		panic(err)
	}

	url := "https://localhost:8080/grpctesting.EchoService/Echo"
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewReader(dat))
	if err != nil {
		panic(err)
	}
	req.Header.Set("TE", "trailers")
	req.Header.Set("Content-Type", "application/grpc")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("response Body:", hex.Dump(body))

}
