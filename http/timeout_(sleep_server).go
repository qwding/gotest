package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// GOOS=linux GOARCH=amd64 go build -o timeout timeout_\(sleep_server\).go

func handler(w http.ResponseWriter, req *http.Request) {
	var err error
	fmt.Println("Come Timeout！")

	fmt.Println("header:", req.Header)

	sleepTimeStr := os.Getenv("SLEEP_TIME")
	sleepTime := 5
	if sleepTimeStr != "" {
		sleepTime, err = strconv.Atoi(sleepTimeStr)
		if err != nil {
			panic("SLEEP_TIME ENV not right. should string")
		}
	}

	time.Sleep(time.Second * time.Duration(sleepTime))

	io.WriteString(w, "返回结果，sleep "+strconv.Itoa(sleepTime)+" 秒.")
}

func visit(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Come visit~")
	io.WriteString(w, "成功visit~")
}

func main() {
	port := ":8008"
	http.HandleFunc("/", handler)
	http.HandleFunc("/visit", visit)
	fmt.Println("Start server at ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println("End Server")
}
