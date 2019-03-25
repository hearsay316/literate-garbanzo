package main

import "fmt"

type aa struct {
	id string
}
type a struct {
	id string
}

func main() {
	ss := aa{"12"}
	s := a{"12"}
	fmt.Println(ss, s)

}
