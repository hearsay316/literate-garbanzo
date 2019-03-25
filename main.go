package main

func main() {
	//var arr [10]int = [10]int{1,2,3,4,5,6,8,7,89,1}
	arr := [11]int{}
	for i, v := range arr {
		println(i, v)
	}
}
