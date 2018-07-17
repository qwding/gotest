package main

import (
	"fmt"
)

var signal chan int

func init() {
	signal = make(chan int, 1)

}
func main() {
	method := "main"
	fmt.Println(method, "before give signal,len is:", len(signal))
	signal <- 1
	fmt.Println(method, " give signal ", 1, "len:", len(signal))
	given()
	fmt.Println(method, "end")
}

func given() {
	method := "given"
	fmt.Println(method, "before give. len:", len(signal))
	/*	_, ok := <-signal
		fmt.Println(method, "val:", "ok:", ok)*/
	if _, ok := <-signal; ok {
		fmt.Println(method, "val:", "ok:", ok, "len:", len(signal))
	} else {
		fmt.Println(method, "val:", "ok:", ok, "len:", len(signal))
	}

	signal <- 2
	fmt.Println(method, "after give", 2, "len:", len(signal))
}
