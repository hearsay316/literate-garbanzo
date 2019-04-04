package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func getHttp(url string) (result string, err error) {
	res, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("http.Get is error", err1)
		err = err1
		return
	}
	defer res.Body.Close()
	for {
		buf := make([]byte, 4096)
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
func SpiderPage2(index int, ch chan int) {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((index-1)*25) + "&filter="
	res, err := getHttp(url)
	if err != nil {
		fmt.Println("getHttp is error", err)
		return
	}
	ret := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	NameAlls := ret.FindAllStringSubmatch(res, -1)

	ret1 := regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	NumAlls := ret1.FindAllStringSubmatch(res, -1)

	ret2 := regexp.MustCompile(`<span>(.*?)人评价</span>`)
	PeopleAlls := ret2.FindAllStringSubmatch(res, -1)
	for _, name := range PeopleAlls {
		fmt.Println("name", name[1])
	}
	SaveFile(index, NameAlls, NumAlls, PeopleAlls)
	ch <- index
}
func SaveFile(index int, NameAlls, NumAlls, PeopleAlls [][]string) {
	path := `F:/coding/gogogo/go正则/` + strconv.Itoa(index) + `页.txt`
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	for i := 0; i < len(NameAlls); i++ {
		_, _ = f.WriteString(NameAlls[i][1] + "\n" + NumAlls[i][1] + "\n" + PeopleAlls[i][1] + "\n")
		_, _ = f.WriteString("_-----------\n")

	}
}
func toWork(start, end int) {
	fmt.Printf("我正在爬取%d到%d页的数据...\n", start, end)
	ch := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage2(i, ch)
	}
	for i := start; i <= end; i++ {
		fmt.Println("爬取完成", <-ch)
	}
}
func main() {
	var start, end int
	fmt.Println("请输入爬虫的起始页")
	fmt.Scan(&start)
	fmt.Println("请输入爬虫的中止页")
	fmt.Scan(&end)
	toWork(start, end)
}
