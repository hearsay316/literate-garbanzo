package main

import "fmt"

func revern(data []int, index int) []int {
	fmt.Println("1", data[index:])
	fmt.Println("1", data[index+1:])

	copy(data[index:], data[index+1:])
	return data[:len(data)-1]
}
func main() {
	arr := []int{1, 23, 456, 78, 9199, 10}
	/*  s1 :=[]int{9,8,7,4,56}
	fmt.Println("si:",cap(s1))
	fmt.Println("arr",cap(arr))
	copy(arr,s1)*/
	//fmt.Println("si:",s1)

	fmt.Println("arr", cap(arr))
	a := revern(arr, 2)
	fmt.Println("a", a)
	fmt.Println("arr", arr)
}
