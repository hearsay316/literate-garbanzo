package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	server, err := net.ResolveUDPAddr("udp", "127.0.0.1:3000")
	if err != nil {
		return
	}
	fmt.Println("UDP", server)
	conn, err := net.ListenUDP("udp", server)
	if err != nil {
		return
	}
	fmt.Println("通信创建", conn)
	defer conn.Close()
	//读取客户端数据
	buf := make([]byte, 4096)
	n, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		return
	}
	fmt.Println("服务器读到数据", string(buf[:n]), addr)
	dayTime := time.Now().String()
	_, err = conn.WriteToUDP([]byte(dayTime), addr)
	if err != nil {
		return
	}
}
