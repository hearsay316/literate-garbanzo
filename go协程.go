package main

import (
	"fmt"
	"time"
)

// 全局 channel
var channel1 = make(chan int)

func print2(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}
func test1() { //先执行
	print2("test1")
	channel1 <- 1024
}
func test2() { //后执行
	print2("test2")
}
func main() {
	go test1()
	var b = "swsw"
	var a = <-channel1
	fmt.Println(a, b)
	go test2()
	for {

	}
}
