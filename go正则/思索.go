package main

import (
	"fmt"
	"go/types"
)

var ch = make(chan int)

func main() {
	var a types.Array
	fmt.Println(b)
	fmt.Println(&a)
	c := &a
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("aaas")
			}
		}
	}()
	fmt.Println("kkk")
	ch <- 1
	fmt.Println("sqwsqw")

}
