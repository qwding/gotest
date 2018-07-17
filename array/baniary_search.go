package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 5, 7, 8, 10, 12, 15, 21, 22}
	target := 10
	if baniary(target, arr) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func baniary(target int, arr []int) bool {
	start := 0
	end := len(arr) - 1

	for end >= start {
		middle := (start + end) / 2
		fmt.Println("debug middle", middle)
		if arr[middle] > target {
			end = middle - 1
		} else if arr[middle] < target {
			start = middle + 1
		} else if arr[middle] == target {
			return true
		}
	}
	return false
}
