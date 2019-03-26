package main

import (
	"fmt"
	"time"
)

func main() {
	myTime := time.NewTimer(time.Second * 3)
	myTime.Reset(1 * time.Second)
	go func() {
		fmt.Println("子go程,定时完毕")
		<-myTime.C
		fmt.Println("子go程,定时完毕")
	}()
	//myTime.Stop()
	for {

	}
}
