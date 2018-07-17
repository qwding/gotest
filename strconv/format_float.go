package main

import (
	"fmt"
	"strconv"
)

func main() {
	// num := 0.00000000056789123

	num := 1234567812345.876545

	res := strconv.FormatFloat(num, 'f', 3, 64)
	fmt.Println("res:", res)
}
