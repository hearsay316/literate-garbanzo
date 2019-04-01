package main

import (
	"fmt"
	"net"
)

// 创建用户结构体

type Client struct {
	C    chan string
	Name string
	Addr string
}

// 创建全局map ,存储在线用户
var onlineMap map[string]Client

//创建全局channel,传递用户数据
var message = make(chan string)

func WriteMsgToClient(client Client, conn net.Conn) {
	for msg := range client.C {
		_, _ = conn.Write([]byte(msg + "\n"))
	}
}
func MakeMsg(clnt Client, msg string) (buf string) {
	buf = "[" + clnt.Addr + "]" + clnt.Name + ":" + msg
	return
}

//noinspection ALL
func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	// 获取网络地址
	netAddr := conn.RemoteAddr().String()
	//  创建新连接用户的结构体 默认是 ip+port
	clnt := Client{make(chan string), netAddr, netAddr}
	onlineMap[netAddr] = clnt
	// 创建专门发送消息的go程
	go WriteMsgToClient(clnt, conn)
	//发送用户上线数据到全局 channel 中
	message <- MakeMsg(clnt, "login")
	// 创建go程
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Printf("检测到客户端:%s退出\n", clnt.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err", err)
				return
			}
			// 数据 保存成sting类型
			msg := string(buf[:n])
			// 将读到的数据广播到message
			message <- MakeMsg(clnt, msg)
		}
	}()
	for {

	}
}

//noinspection ALL
func Manager() {
	// 初始化map
	onlineMap = make(map[string]Client)
	// 循环读通道channel
	for { //循环读取 用户
		msg := <-message
		//循环发送消息
		for _, clnt := range onlineMap {
			clnt.C <- msg
		}
	}

}
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("listener is err", err)
		return
	}
	defer listener.Close()
	// 创建管理者go程
	go Manager()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn is error", err)
			return
		}
		// 启动gc程 处理请求
		go HandlerConnect(conn)

	}
}
