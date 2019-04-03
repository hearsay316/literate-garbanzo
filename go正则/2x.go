package main

import (
	"fmt"
	"regexp"
)

func main() {
	//str:="3.14 123 123.321 .68 haha 1.0 abc 7. ab.3 66.6 123."
	str := "ppp12bbbbb12.31ccc"
	ret := regexp.MustCompile(`[0-9][0-9]|[0-9]+\.[0-9]`)
	all := ret.FindAllStringSubmatch(str, -1)

	fmt.Println(all)
}
