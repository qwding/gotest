package main

import (
	"fmt"
	"reflect"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	aba := live()
	if aba == nil {
		// fmt.Println("live:", live())
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("aba:", aba)
		fmt.Println("type:", reflect.TypeOf(aba))
		fmt.Println("BBBBBBB")
	}
}
