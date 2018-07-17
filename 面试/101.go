package main

import (
	"fmt"
)

func main() {
	x := make(map[int]int)
	for i := 0; i < 30; i++ {
		x[i] = i
	}
	for k, v := range x {
		fmt.Println(k, v)
	}
}

// func main() {
// 	x := make(map[int]int)
// 	x[1] = 11
// 	x[3] = 33
// 	x[2] = 22

// 	for k, v := range x {
// 		fmt.Println(k, v)
// 	}
// }
