package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start ")
	go func() {
		time.Sleep(time.Second * 100)
	}()

	select {}
}
