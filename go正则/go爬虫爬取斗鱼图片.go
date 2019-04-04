package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SpiderPage(i int, page chan int) {
	url := "https://www.pengfue.com/xiaohua_" + strconv.Itoa(i) + ".html"
	result, err := httpget(url)
	if err != nil {
		fmt.Println("错误->>http get ")
		return
	}
	ret := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	all := ret.FindAllStringSubmatch(result, -1)
	flieTitle := make([]string, 0)
	fileContent := make([]string, 0)
	for _, jokeURL := range all {
		fmt.Println(jokeURL[1])
		title, content, err := SpidJokePage(jokeURL[1])
		if err != nil {
			fmt.Println("SpidJokePage error  ")
			continue
		}
		flieTitle = append(flieTitle, title)
		fileContent = append(fileContent, content)
	}
	fileCreate(i, flieTitle, fileContent)
	page <- i
}
func fileCreate(i int, flieTitle, fileContent []string) {
	path := "F:/coding/gogogo/go正则/" + strconv.Itoa(i) + "页.txt"
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Create is err", err)
		return
	}
	defer f.Close()
	for i := 0; i < len(flieTitle); i++ {
		_, _ = f.WriteString(flieTitle[i] + "\r\n" + fileContent[i] + "\n")
		_, _ = f.WriteString("--------------------------------\n")
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

func toWork(start int, eng int) {
	fmt.Printf("正在爬取%d到%d的起始页:", start, eng)
	page := make(chan int)
	for i := start; i <= eng; i++ {
		go SpiderPage(i, page)
	}
	for i := start; i <= eng; i++ {
		fmt.Printf("这个%d是并发\n", <-page)
	}
}
func main() {
	url := "https://www.douyu.com/g_yz"
	resp, err := httpget(url)
	if err != nil {
		fmt.Println("httpget is error")
		return
	}
	ret := regexp.MustCompile(`<img src="(?s:(.*?))" class="DyImg-content is-normal"`)
	alls := ret.FindAllStringSubmatch(resp, -1)
	for _, imgURl := range alls {
		fmt.Println("img", imgURl[1])
	}
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
