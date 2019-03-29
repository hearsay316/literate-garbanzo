package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 3, 4, 56, 7}
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", &arr)
}
