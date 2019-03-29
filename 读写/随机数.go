package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//根据时间衍生
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(100))
	}
}
