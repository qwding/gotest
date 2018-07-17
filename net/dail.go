package main

import (
	"fmt"
	"net"
)

func main() {
	var url = ""
	for url != "0" {
		fmt.Scanln(&url)
		_, err := net.Dial("tcp", url)
		if err != nil {
			fmt.Println("can't connect it ", url, " ", err)
		} else {
			fmt.Println("connect to url ", url)
		}
	}
}
