package main

import (
	"fmt"
	// "os"
	"time"
)

var c = make(chan int)

func main() {
	go Done()
	go Give()
	select {}
}

/*
结论：
	1. select 也可以用break，在select break只会跳出select，不会跳出for，所以要用tag来指定跳出循环
	2. select计时时候，如果令起一个线程计时，那么可以在计时内无限接收传来消息了，适应需求
	3. 一般时候生产者负责关闭chan，消费者不负责关闭，因为这样容易panic “send on closed channel”
*/

func Done() {
	// defer close(c)
	fmt.Println("start done")
	timeout := make(chan bool, 1)
	defer close(timeout)
	go func() {
		time.Sleep(time.Second * 6)
		timeout <- true
	}()

Loop:
	for {
		select {
		case tmp := <-c:
			fmt.Println("tmp:", tmp)
			// close(timeout)
			return
		case <-timeout:
			fmt.Println("timeout true")
			break Loop
		}
	}
	fmt.Println("success break loop")
}

func Give() {
	fmt.Println("in given")
	a := 0
	// for {
	time.Sleep(time.Second * 2)
	// fmt.Println("in given give a:", a)
	c <- a
	a++
	// }
}
