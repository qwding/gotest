package main

import (
// "fmt"
)

func main() {
	arr1 := []string{"qqq", "www", "eeee"}
	arr2 := []string{"zzz", "xxxx", "cccc"}
	arr1 = append(arr1, arr2...)

}
