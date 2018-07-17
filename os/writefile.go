package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//writestring会 重新写文件，覆盖之前写 东西
	file, _ := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0666)
	// file.WriteString("1aaaa111")
	// file.WriteString("2bbbb2222")
	// file.WriteString("3cccc333")
	defer file.Close()
	// fmt.Println("write over")

	// d2 := []byte{115, 111, 109, 101, 10}
	// n2, _ := file.Write(d2)
	// fmt.Println("wrote %d bytes\n", n2)

	// file.Sync()

	w := bufio.NewWriter(file)
	n4, _ := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
