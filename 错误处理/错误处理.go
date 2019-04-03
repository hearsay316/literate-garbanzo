package main

import (
	"fmt"
	"net"
	"os"
)

func errFunc(err error, info string) {
	if err != nil {
		fmt.Println("info", info, "err", err)
		os.Exit(1) // 将当先进程结束
	}
}
func main() {
	listtener, err := net.Listen("tcp", "127.0.0.1:8001")
	errFunc(err, "listener is err")
	defer listtener.Close()
	conn, err := listtener.Accept()
	errFunc(err, "listener is err")
	defer conn.Close()
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	errFunc(err, "conn.Read")
	fmt.Println(string(buf[:n]))
}
