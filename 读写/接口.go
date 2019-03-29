package main

import "fmt"

type serApper interface {
	SayHello()
}
type person struct {
	name string
	sex  string
	age  int
}
type student struct {
	person
	score int
}
type teacher struct {
	person
	subject string
}

func (s *student) SayHello() {
	fmt.Printf("大家好,我是%s,我是%s,我今年%d岁了,我的成绩是%d分\n", s.name, s.sex, s.age, s.score)
}
func (t *teacher) SayHello() {
	fmt.Printf("大家好,我是%s,我是%s,我今年%d岁了,我的教授是%s", t.name, t.sex, t.age, t.subject)
}
func main() {
	var stu student = student{person{"小明", "男", 28}, 98}
	var tea teacher = teacher{person{"法师", "男", 32}, "语文"}
	/*stu.SayHello()
	  tea.SayHello()*/
	var h serApper
	var j serApper
	j = &tea
	h = &stu
	h.SayHello()
	j.SayHello()
}
