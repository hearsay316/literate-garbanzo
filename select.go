package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true
		runtime.Goexit()
	}()
	for {
		select {
		case num := <-ch:
			fmt.Println("读到", num)
		case <-quit:
			//break // 不能终止for ,只能跳出select
			return //终止进程
		}
	}
}
