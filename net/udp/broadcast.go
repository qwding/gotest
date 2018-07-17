package main

import (
	"net"
)

func main() {
	// 这里设置发送者的IP地址，自己查看一下自己的IP自行设定
	laddr := net.UDPAddr{
		IP:   net.IPv4(192, 168, 56, 1),
		Port: 3000,
	}
	// 这里设置接收者的IP地址为广播地址
	raddr := net.UDPAddr{
		// IP:   net.IPv4(192, 168, 56, 255),
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 3000,
	}
	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		println(err.Error())
		return
	}
	conn.Write([]byte(`hello peersthis is broadcast test.`))
	conn.Close()
}
