package main

import (
	"fmt"
	"runtime"
)

func main() {
	arch := runtime.GOARCH
	os := runtime.GOOS
	fmt.Println("arch os", arch, os)
	numCpu := runtime.NumCPU()
	fmt.Println("cpu is:", numCpu)
}
