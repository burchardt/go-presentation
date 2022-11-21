package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3, 6)
	s[0], s[1], s[2] = "a", "b", "c"
	fmt.Println("s:", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("s:", s)

	l := s[1:3]
	fmt.Println("l:", l)
	l = append(l, "x", "y")
	fmt.Println("s:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c:", c)
}
