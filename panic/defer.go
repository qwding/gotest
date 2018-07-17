package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	var i int = 0
	defer func() {
		fmt.Println("come into defer")
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for s := range c {
			fmt.Printf("received ctrl+c(%v)\n", s)
			os.Exit(0)
		}
	}()

	for {
		fmt.Println("output ", i)
		i++
		/*if i == 3 {
			panic("123")
		}*/
		time.Sleep(time.Second * 2)
	}
}
