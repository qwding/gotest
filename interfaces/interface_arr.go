package main

import (
	"fmt"
)

func main() {
	Done()
}

func testfunc(args ...interface{}) {
	fmt.Println("args:", args, "len(args):", len(args))
}

func Done() {
	// arr := []string{"abc", "bcd", "xyz"}
	// testfunc(arr...)
}
