package main

import (
	"fmt"
	"time"
)

var c1 chan int

func init() {
	c1 = make(chan int)
}

func main() {
	// go gen()
	fmt.Println("xo")
	c1 <- 234
	time.Sleep(time.Second * 3)
	fmt.Println("after sleep")
	// fmt.Println("first get:", <-c1)
	close(c1)
	fmt.Println("after close get:", <-c1)
	// time.Sleep(time.Second * 10)
}

func gen() {
	fmt.Println("before given:")
	c1 <- 123
	fmt.Println("has given")
}
