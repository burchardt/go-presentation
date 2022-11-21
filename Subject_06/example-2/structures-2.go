package main

import "fmt"

type Person struct {
	Name string
	Age  uint8
}

type Worker struct {
	Person
	Salary uint32
}

func riseSalary(w *Worker, by uint32) {
	w.Salary += by
}

func printWorker(w *Worker) {
	fmt.Printf("{Name: %s, Age: %d, Salary: %d}\n", w.Name, w.Age, w.Salary)
}

func main() {
	w1 := &Worker{Person{"Bob", 25}, 2345}
	printWorker(w1)

	w2 := &Worker{Person{Age: 30, Name: "Alice"}, 3456}
	printWorker(w2)

	riseSalary(w1, 1000)
	printWorker(w1)
}
