package main

import (
	// "crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "golang.org/x/net/http2"
)

func main() {
	// hc := &http.Client{
	// 	Transport: &http2.Transport{
	// 		TLSClientConfig: &tls.Config{
	// 			InsecureSkipVerify: true,
	// 		},
	// 	},
	// }

	hc := &http.Client{}

	resp, err := hc.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", body)

}
