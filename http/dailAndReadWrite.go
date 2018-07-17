package main

import (
	// "bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	//mysql-0ne48.tenxcloud.net:40982
	//mysql-vcihs.tenxcloud.net:40967
	var url string
	for url != "0" {
		fmt.Println("")
		duration, errDuration := time.ParseDuration("30s")
		// fmt.Println("duration is ", duration)
		if errDuration != nil {
			fmt.Println("errDuration is ", duration, errDuration)
		}
		fmt.Println("please give your connect url")
		fmt.Scanln(&url)
		// conn, err := net.Dial("tcp", url)
		conn, err := net.DialTimeout("tcp", url, duration)
		if err != nil {
			fmt.Println("can't connect it ", url, " ", err)
		} else {
			fmt.Println("connect to url ", url)
			fmt.Printf("conn is %#v\n", conn)
			resWrite, errWrite := conn.Write([]byte("ok"))
			if errWrite != nil {
				fmt.Println("errwrite", errWrite)
			} else {
				fmt.Println("write string is ", resWrite)
			}
			fmt.Println("Start read server")
			res, errRead := conn.Read([]byte("ok"))
			if errRead != nil {
				fmt.Println("errRead error", errRead)
			} else {
				fmt.Printf("conn neiron \n", res)
			}
		}
	}
}
