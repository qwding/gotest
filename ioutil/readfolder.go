package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	skillfolder := `d:\testcode`
	// 获取所有文件
	files, _ := ioutil.ReadDir(skillfolder)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name() + "/")
		} else {
			fmt.Println(file.Name())
		}
	}
}
