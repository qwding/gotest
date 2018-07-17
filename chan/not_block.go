package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	// 非阻塞读取
	select {
	case <-c:
	default:
	}
	fmt.Println("non-blocking")
	// 非阻塞写
	select {
	case c <- 1:
	default:
	}
	fmt.Println("non-blocking")

}
