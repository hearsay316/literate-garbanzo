package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc a7c mfc cat 8ca azc aba"
	// 解析编译正则
	res := regexp.MustCompile(`a.c`)
	alt := res.FindAllStringSubmatch(str, -1)
	fmt.Println(alt)
}
