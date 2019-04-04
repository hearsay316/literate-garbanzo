package main

import "fmt"

func add(b *[5]int) {
	b[2] = 5
}
func add1(b *[]int) {
	*b = append(*b, 5)
}
func main() {
	var a []int
	var b [5]int
	a = []int{123, 456, 78}
	add(&b)
	add1(&a)
	fmt.Println(a, b)
}
