package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	tf_bak := "test.txt"
	fbak, err := os.OpenFile(tf_bak, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer fbak.Close()

	bw := bufio.NewWriter(fbak)

	bw.Write([]byte("xyz"))
	bw.Flush()

	bw.Write([]byte("opq"))
	bw.Flush()
}
