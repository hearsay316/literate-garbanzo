package main

import (
	"fmt"
	"net/http"
)

func hanler(w http.ResponseWriter, r *http.Request) {
	n, err := w.Write([]byte("hello 9fud9fd9"))
	if err != nil {
		return
	}
	fmt.Println(n)
}
func main() {
	// 注册回调函数
	http.HandleFunc("/itcast", hanler)
	err := http.ListenAndServe("127.0.0.1:8001", nil)
	if err != nil {
		return
	}
}
