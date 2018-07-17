package main

import (
	"fmt"
	"time"
)

var signal chan int

func main() {
	method := "main"
	fmt.Println(method, "start")
	signal = make(chan int, 1)

	// go given()
	// go got()
	go while()

	// time.Sleep(time.Second * 10)
	select {}
	fmt.Println(method, "end")
}

func given() {
	method := "given"
	for {
		fmt.Println(method, "before give.")
		time.Sleep(time.Second * 3)

		signal <- 1
		fmt.Println(method, "after give")
	}
}

func got() {
	method := "got"
	for {
		fmt.Println(method, "before get data.")
		<-signal
		fmt.Println(method, "after get data.")
	}
}

func while() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("while")
	}
}
