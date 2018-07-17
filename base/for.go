package main

import (
	"fmt"
)

func main() {

	for j := 0; j < 5; j++ {
	J:
		for i := 0; i < 10; i++ {
			if i > 6 {

				break J //现在终止的是j 循环，而不是i的那个
			}
			fmt.Println(i)
		}
		fmt.Println("a")
	}
}
