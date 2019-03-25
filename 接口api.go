package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("./2.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	fp.WriteString("这是中文")
	fp.WriteString("\n这是中文")
	fmt.Println(*fp)
	defer fp.Close()
}
