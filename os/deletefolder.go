package main

import (
	"fmt"
	"os"
)

func main() {
	temp, _ := os.Getwd()
	errfoder := os.RemoveAll(temp + "/doc")
	fmt.Println("remvoerfolder", errfoder)
	errFile := os.Remove(temp + "/Docker.zip")
	fmt.Println("removefile", errFile)
	fmt.Println("over")
}
