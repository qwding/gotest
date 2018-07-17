package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("error conn,", err)
			continue //改成break试试
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())

		fmt.Printf("tcp conn %#v\n", tcpConn)

		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(string(message))
		msg := time.Now().String() + "  " + string(message)
		b := []byte(msg)
		conn.Write(b)
		if string(message) == "exit\n" {
			conn.Close()
			return
		}
	}
}
