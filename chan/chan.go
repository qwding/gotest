package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("come int")
	ca := make([]chan string, 10)
	for i := 0; i < 10; i++ {
		ca[i] = make(chan string)
		go route(ca[i])

	}
	for i := 0; i < 10; i++ {
		fmt.Println("ca i", <-ca[i])
	}
	// c := make(chan string)
	fmt.Println("come out :", "<-c")
	time.Sleep(time.Second * 2)
}

func route(c chan string) {
	c <- "xyz"
}
