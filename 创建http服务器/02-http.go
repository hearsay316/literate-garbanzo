package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func OpenSennFile(fname string, w http.ResponseWriter) {
	decodeurl, err := url.QueryUnescape(fname)
	if err != nil {
		fmt.Println(err)
	}
	pathName := "F:/BaiduYunDownload/第二阶段/2-8 HTTP服务器开发/ 2 HTTP服务器实现" + decodeurl
	fmt.Println(pathName)
	f, err := os.Open(pathName)
	if err != nil {
		fmt.Println("Open is error")
		w.Write([]byte("错误"))
		return
	}
	buf := make([]byte, 4096)
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			return
		}
		w.Write(buf[:n])
	}
}

func myHanle(w http.ResponseWriter, r *http.Request) {
	OpenSennFile(r.URL.String(), w)
	fmt.Println("header", r.Header)
	fmt.Println("URL", r.URL)
	fmt.Println("Host", r.Host)
	fmt.Println("remoteAddr", r.RemoteAddr)
	fmt.Println("Body", r.Body)
}
func main() {
	http.HandleFunc("/", myHanle)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		return
	}
}
