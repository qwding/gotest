package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
	}

	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second * 2)
			wg.Done()

		}()
	}
	fmt.Println("exit")
	wg.Wait()
	fmt.Println("after wait.")
}

func add(wg sync.WaitGroup) {
	wg.Add(1)
}

func done(wg sync.WaitGroup) {
	wg.Done()
}
