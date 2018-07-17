package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file1, err1 := os.OpenFile("testFile.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err1 != nil {
		log.Println("err1", err1)
	}
}
