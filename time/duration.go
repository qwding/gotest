package main

import (
	"fmt"
	"time"
)

func main() {
	duration, err := time.ParseDuration("-5m")
	fmt.Println("duration ", duration, err)

	timenow := time.Now()
	fmt.Println("timenow before", timenow)
	timeadd := timenow.Add(duration)
	fmt.Println("time ", timeadd)
}
