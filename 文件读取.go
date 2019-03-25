package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("F:/BaiduYunDownload/第二阶段/2-1 GO语法加强/5 Go语言与文件IO操作/01-复习.mp4")
	if err != nil {
		fmt.Println("文件不存在")
	}
	defer f.Close()
	// 创建文件
	c, err := os.Create("F:/BaiduYunDownload/第二阶段/2-1 GO语法加强/5 Go语言与文件IO操作/01-复习1.mp4")
	if err != nil {
		fmt.Println("文件创建失败")
	}
	defer c.Close()
	buf := make([]byte, 74096)
	for {
		n, err := f.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("读取结束", n)
			return
		}
		w, err := c.Write(buf[:n])
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println(w)
	}

}
