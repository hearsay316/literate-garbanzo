package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "死神来了, 死神 bye bye"
	b := a[12:]
	pos := strings.Index(a[12:], "死神")
	fmt.Print(len(a), pos, b)
}
