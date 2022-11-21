package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

func printPerson(p Person) {
	fmt.Printf("{Name: %s, Age: %d}\n", p.Name, p.Age)
}

func getOlder(p *Person, by uint8) {
	p.Age += by
}

func main() {
	p1 := Person{"Bob", 20}
	printPerson(p1)

	p2 := Person{Age: 25, Name: "Alice"}
	printPerson(p2)

	getOlder(&p2, 10)
	printPerson(p2)
}
