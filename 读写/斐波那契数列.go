package main

import (
	"fmt"
	"runtime"
)

func num(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num1 := <-ch:
			fmt.Println(num1, "")
		case <-quit:
			runtime.Goexit()
		}
	}
}
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go num(ch, quit)
	x, y := 1, 1
	for i := 0; i < 20; i++ {
		ch <- x
		x, y = y, x+y

	}
	quit <- true
}
