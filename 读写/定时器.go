package main

import (
	"fmt"
	"time"
)

/*func main01(){
	fmt.Println(time.Now(),"/n")
	myTime:=time.NewTimer(time.Second*2)
	nowTime:=<-myTime.C
	fmt.Println(time.Now(),"/n",nowTime)
}*/
func main() {
	//定时器
	//time.Sleep(time.Second*20)
	// 02
	fmt.Println(time.Now())
	a := <-time.After(time.Second * 3)
	fmt.Println(a)
}
