package main

import (
	"fmt"
	"time"
)

func main() {

	timenow := time.Now()
	fmt.Println("time before:", timenow)

	timeAfter := timenow.UTC()
	fmt.Println("time after :", timeAfter)
}
