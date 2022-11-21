package main

import "fmt"
import "math"

type rectangle struct {
	width, height float64
}

func (r rectangle) surface() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (r rectangle) toString() string {
	return fmt.Sprintf("rectangle{width: %.2f, heigh:%.2f}", r.width, r.height)
}

type circle struct {
	radius float64
}

func (c circle) surface() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) toString() string {
	return fmt.Sprintf("circle{radius: %.2f}", c.radius)
}

type geometry interface {
	surface() float64
	perimeter() float64
	toString() string
}

func measure(g geometry) {
	fmt.Println(g.toString())
	fmt.Printf("Surface: %f, Perimeter: %f\n", g.surface(), g.perimeter())
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
