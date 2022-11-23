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

	var big uint64 = 0
	fmt.Println(big)
	
	s := "short"
	fmt.Println(s)

	global = 100
	fmt.Println(global)
}
