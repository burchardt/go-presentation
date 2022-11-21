package main

import (
	"fmt"
)

const g = 9.81

const (
	a = 1
	b = true
	c = "constant"
)

type Season int

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

func main() {
	fmt.Println(a, b, c)

	const x = a + 2
	fmt.Println(x)

	const season = Summer
	fmt.Print(season)
}
