package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	err := ioutil.WriteFile("test/test.txt", []byte("&&&&&&&&&&&&&&&"), 0644)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("finish")
}
