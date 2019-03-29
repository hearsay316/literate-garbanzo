package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}
type student struct {
	*person
	id    int
	score int
}

func main() {
	var stu student
	//开创内存
	stu.person = new(person)
	stu.name = "郭襄"
	stu.person.name = "郭小姐"
	stu.id = 101
	stu.age = 78
	stu.sex = "女"
	fmt.Println(stu)
}
