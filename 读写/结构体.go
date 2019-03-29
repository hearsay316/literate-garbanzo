package main

import "fmt"

type app struct {
	id   int
	name string
	sex  string
	age  int
}
type stud struct {
	id    int
	name  string
	score [3]int
}

func test(s app) {
	fmt.Println(s)
}
func main() {
	f := app{id: 01, name: "张立群", sex: "男", age: 23}
	f2 := app{01, f.name, f.sex, f.age}
	test(f2)
	arr := []stud{
		{101, "小明", [3]int{101, 100, 100}},
		{102, "小红", [3]int{99, 98, 97}},
	}
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := 0; j < len(arr[i].score); j++ {
			sum += arr[i].score[j]
		}
		fmt.Printf("第%d学生总成绩:%d 平均成绩: %d\n", i+1, sum, sum/len(arr[i].score))
	}
}
