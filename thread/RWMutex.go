package main

import (
	"fmt"
	// "log"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var m *sync.RWMutex

var arr []string = []string{"x1", "x2", "x3", "x4"}

/*
*经测试，发现RWMutex Lock后不能读，不能写。
 */

func main() {
	runtime.GOMAXPROCS(2)

	m = new(sync.RWMutex)
	/*	go read(1)
		go read(2)
		go read(3)
	*/
	go funcVisit(1)
	go funcVisit(2)
	go visit(3)
	for {
		// fmt.Println("in for")
	}
}

func read(num int) {
	fmt.Println("come into", num)
	m.Lock()
	arr = append(arr, strconv.Itoa(num))                          //can't write,
	fmt.Println("process ", num, "arr ", arr, "time", time.Now()) //can't read
	time.Sleep(time.Second * 2)
	m.Unlock()
}

func funcVisit(num int) {
	fmt.Println("come into ", num, "time", time.Now())
	m.Lock()
	visit(num)
	time.Sleep(time.Second * 2)
	m.Unlock()
}

func visit(i int) {
	fmt.Println("visit before", i, arr, "time", time.Now())
	time.Sleep(time.Second * 3)

	fmt.Println("visit after", i, arr, "time", time.Now())
}
