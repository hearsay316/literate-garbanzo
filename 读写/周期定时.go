package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool)
	fmt.Println(time.Now())
	myTimer := time.NewTicker(time.Second)
	i := 0
	go func() {
		for {
			nowTime := <-myTimer.C
			i++
			fmt.Println(nowTime)
			if i == 8 {
				quit <- true
			}
		}
	}()
	<-quit
}
