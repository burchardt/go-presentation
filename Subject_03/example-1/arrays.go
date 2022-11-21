package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 100
	a[1] = 200
	fmt.Println("a:", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	c := b[1:3]
	fmt.Println("c:", c)
}
