package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.OpenFile("test2.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("err:", err)
	}
	defer f.Close()

	// bw := bufio.NewReadWriter(f, f)

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadBytes(byte('\n'))
		if err != nil {
			if err == io.EOF {
				return
			}
			return
		}
		_ = line
	}

	fmt.Println("end")
}
