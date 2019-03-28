package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr()
	fmt.Println(addr, "地址是")
	buf := make([]byte, 4096)
	// 循环读
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println("读到的数据", string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
func main() {
	comm, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		return
	}
	defer comm.Close()
	for {
		fmt.Println("服务器启动成功")
		conn, err := comm.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		HandlerConnect(conn)
	}
}
