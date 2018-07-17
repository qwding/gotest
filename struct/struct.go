package main

import "fmt"

type Test struct {
	Name   string
	Age    int
	School string
}

func main() {
	test1 := &Test{Name: "ding", Age: 16, School: "hit"}
	test2 := &Test{Name: "yao", Age: 17, School: "hit"}
	test3 := &Test{Name: "ding", Age: 16, School: "hit"}
	//两个都是false，因为test1,test2,test3都是指针啊。他们地址不一样
	fmt.Println("test1 == test2", test1 == test2)
	fmt.Println("test1 == test3", test1 == test3)

	//后一个是true，是值了
	fmt.Println("test1 == test2", *test1 == *test2)
	fmt.Println("test1 == test3", *test1 == *test3)
}
