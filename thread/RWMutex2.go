package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex

func main() {
	m = new(sync.RWMutex)

	// 写的时候啥也不能干
	go read(1)
	go read(2)
	// time.Sleep(time.Second * 1)
	go write(3)

	time.Sleep(5 * time.Second)
}

func read(i int) {
	fmt.Println(i, "read start", time.Now())

	m.RLock()
	fmt.Println(i, "reading", time.Now())
	time.Sleep(2 * time.Second)
	m.RUnlock()

	fmt.Println(i, "read over", time.Now())
}

func write(i int) {
	fmt.Println(i, "write start", time.Now())

	m.Lock()
	fmt.Println(i, "writing", time.Now())
	time.Sleep(2 * time.Second)
	m.Unlock()

	fmt.Println(i, "write over", time.Now())
}
