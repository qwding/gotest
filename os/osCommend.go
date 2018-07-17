package main

import (
	"fmt"
)

func main() {
	str1 := "good"
	str2 := "boy"
	abc := fmt.Sprintln("ding", str1, str2)
	fmt.Println(abc)

	qwe := "ding" + str1 + str2
	fmt.Println(qwe)
}
