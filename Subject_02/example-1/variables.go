package main

import (
	"fmt"
)

var global int

func main() {
	var a int
	fmt.Println(a)

	var x, y = 10, 20
	fmt.Println(x, y)

	s := "short"
	fmt.Println(s)

	global = 100
	fmt.Println(global)
}
