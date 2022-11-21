package main

import "fmt"

func incIntBy(iPtr *int, by int) {
	*iPtr += by
}

// The function has no effect !!!
func setZero(i int) {
	i = 0
}

func main() {
	x := 10
	incIntBy(&x, 2)
	setZero(x)
	fmt.Println(x)
}
