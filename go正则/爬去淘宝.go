package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://detail.tmall.com/item.htm?id=586524097919&ali_refid=a3_430652_1007:1104985326:N:50010402_0_100:b2bb12503d1d0235ff77588ca87ab08a&ali_trackid=1_b2bb12503d1d0235ff77588ca87ab08a&spm=a231k.7971385.300002.1")
	if err != nil {
		return
	}
	var result string
	defer res.Body.Close()
	for {
		buf := make([]byte, 4096)
		n, err := res.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			return
		}
		result += string(buf[:n])
	}

	fmt.Println(result)
}
