package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定地址和协议和ip+端口号
	listen, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		return
	}
	//关闭
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端成功连接")
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	_, _ = conn.Write(buf[:])
	fmt.Println(string(buf[:n]), "服务器错误")

}
