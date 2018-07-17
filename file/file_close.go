package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	Done()
	// GetFileEnv()
}

// 测试 重复调用file.Cloes()是否报错
func Done() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file,", err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println("error123:", err)
	}
	err = file.Close()
	if err != nil {
		fmt.Println("error345:", err)
	}
	os.Remove("test.txt")
}

func GetFileEnv() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file,", err)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("content,", string(content))

}
