package main

import (
	"fmt"
	"os"
	"time"
)

var fileName = "test_log.log"

func main() {
	for {
		writeLog()
		time.Sleep(time.Second * 1)
	}
}

func writeLog() {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer f.Close()

	date := time.Now().String()
	f.WriteString(date + "qidaye")
	f.Write([]byte{byte('\n')})
}
