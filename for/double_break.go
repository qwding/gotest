package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
	LABEL:
		for j := 0; j < 100; j++ {
			fmt.Println("i:", i, j)
			if j == 4 {
				break LABEL
			}
		}
	}
}
