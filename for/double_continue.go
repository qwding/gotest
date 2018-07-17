package main

import (
	"fmt"
)

func main() {

Lable:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("i", i, "j:", j)
			if j == 3 {
				fmt.Println("continue:")
				continue Lable
			}
		}
	}
}
