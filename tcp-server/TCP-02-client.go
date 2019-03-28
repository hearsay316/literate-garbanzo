package main

import (
	"fmt"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		return
	}
	defer dial.Close()
	_, _ = dial.Write([]byte("Are you Ready?"))
	buf := make([]byte, 4096)
	n, err := dial.Read(buf)
	if err != nil {
		return
	}
	fmt.Println("这个是服务器回发", string(buf[:n]))
}
