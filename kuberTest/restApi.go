package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://proxy2.tenxcloud.net:40035"
	uri := "/api/v1beta3/nodes"
	method := "GET"

	request, err := http.NewRequest(method, url+uri, nil)
	if err != nil {
		fmt.Println("err ", err)
	}
	auth := "DPI4ipdNq4aGsccFF84wxEb0CjafwPid"
	request.Header.Set("authorization", "bearer "+auth)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("err ", err)
	}

	body := resp.Body
	defer body.Close()
	bodybyte, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println("body is ", string(bodybyte))

}
