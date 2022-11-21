package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println("m:", m)

	delete(m, "two")
	fmt.Println("m:", m)

	i := m["one"]
	fmt.Println("i:", i)

	j := m["seven"]
	fmt.Println("j:", j)

	k, ok := m["five"]
	fmt.Println("k, ok:", k, ok)

	n := map[string]int{"x": 101, "y": 102, "z": 103}
	fmt.Println("n:", n)
}