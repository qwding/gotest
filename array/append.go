package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 5)
	fmt.Println("arr:", arr)
	arr = append(arr, 4)
	fmt.Println("arr:", arr)

	arr2 := make([]int, 0, 5)
	fmt.Println("arr:", arr2)
	arr2 = append(arr2, 4)
	fmt.Println("arr:", arr2)

}
