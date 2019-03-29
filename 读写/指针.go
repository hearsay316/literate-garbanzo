package main

import "fmt"

func main() {
	var p *string
	fmt.Println(p)
	p = new(string)
	fmt.Println(p)
	fmt.Println(*p, 123)
	if nil == p {
		fmt.Println(*p, 123)
	} else {
		fmt.Println(*p, 321)
	}

}
