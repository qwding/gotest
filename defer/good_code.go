package main

import (
	"fmt"
)

func print(pi *int) { fmt.Println(*pi) }

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i is:", &i)
		defer fmt.Println(i)
		defer func() { fmt.Println(i) }()
		defer func(i int) { fmt.Println(i) }(i)
		defer print(&i)
		go fmt.Println(i)
		go func() { fmt.Println(i) }()
	}
}

/*
	defer fmt.Println(i)   // 0~9
	defer func() { fmt.Println(i) }()  // 10 个10
	defer func(i int) { fmt.Println(i) }(i)	//0~9


*/
