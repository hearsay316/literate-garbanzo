package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	s := arr[1:5:7]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s2 := s[0:6]
	fmt.Println(s2)
	s2[5] = 12
	fmt.Println(1, s2)
	fmt.Println(s)
	fmt.Println(arr)
}
