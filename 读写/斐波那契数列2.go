package main

import (
	"fmt"
	"runtime"
)

func mun(ch chan int, quit chan bool) {
	for {
		select {
		case nu := <-ch:
			fmt.Println("nu", nu)
		case <-quit:
			runtime.Goexit()
		}
	}
}
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go mun(ch, quit)
	x, y := 1, 1
	for i := 0; i < 30; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}
