package main

import "fmt"

func main() {
	arr := [10]int{1, 8, 9, 5, 1, 23, 6, 4, 455, 15}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
