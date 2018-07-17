package main

import (
	"fmt"
	"time"
)

var lastPrint time.Time

func init() {
	lastPrint = time.Now()
}

func main() {
	for {
		if time.Now().Sub(lastPrint) > time.Second*time.Duration(2) {
			printxyz()
		}
	}
}
func printxyz() {
	lastPrint = time.Now()
	fmt.Println("print xyz", time.Now())
}
