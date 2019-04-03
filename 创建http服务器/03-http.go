package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.baidu.com")
	defer res.Body.Close()
	fmt.Println("res", res.Header)
	buf := make([]byte, 4096)
	var result string
	for {
		n, err := res.Body.Read(buf)
		if err != nil {
			break
		}
		if n == 0 {
			fmt.Println("读完了")
			return
		}
		result += string(buf[:n])
	}
	fmt.Println(result)
}
