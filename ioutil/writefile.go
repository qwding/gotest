package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	err := ioutil.WriteFile("test.txt", []byte("xyz111222"), 0700)
	fmt.Println("finish", err)
}
