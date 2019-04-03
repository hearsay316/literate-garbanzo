package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func SpiderPage(i int) {
	url := "https://www.pengfue.com/xiaohua_" + strconv.Itoa(i) + ".html"
	result, err := httpget(url)
	if err != nil {
		fmt.Println("错误->>http get ")
		return
	}
	ret := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	all := ret.FindAllStringSubmatch(result, -1)

	for _, jokeURL := range all {
		fmt.Println(jokeURL[1])
		title, content, err := SpidJokePage(jokeURL[1])
		if err != nil {
			fmt.Println("SpidJokePage error  ")
			continue
		}
		fmt.Println(title, content)
	}

}
func SpidJokePage(url string) (title, content string, err error) {
	result, err1 := httpget(url)
	if err1 != nil {
		err = err1
		return
	}
	relt := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	all := relt.FindAllStringSubmatch(result, 1)
	for _, tempTitle := range all {
		title = tempTitle[1]
		title = strings.Replace(title, " ", "", -1)
		title = strings.Replace(title, "\n", "", -1)
		break
	}
	relt2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href`)
	all2 := relt2.FindAllStringSubmatch(result, 1)
	for _, tempContent := range all2 {
		content = tempContent[1]
		content = strings.Replace(content, " ", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		break
	}
	return

}
func httpget(url string) (result string, err error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("错误->>http get ")
		return
	}
	defer res.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := res.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}
func toWork(start int, eng int) {
	fmt.Printf("正在爬取%d到%d的起始页:", start, eng)
	for i := start; i < eng; i++ {
		SpiderPage(i)
	}
}
func main() {
	var start, eng int
	fmt.Println("请输入爬取的起始点(>1):")
	fmt.Scan(&start)
	fmt.Println("请输入终止页(>start):")
	fmt.Scan(&eng)
	toWork(start, eng)
}
