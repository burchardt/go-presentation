package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r rectangle) surface() int {
	return r.width * r.height
}

func (r rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rectangle{width: 8, height: 3}

	fmt.Println("surface: ", r.surface())
	fmt.Println("perimeter:", r.perimeter())
}
