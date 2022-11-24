package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	switch i {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Printf("Other number: %d\n", i)
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!!!")
	default:
		fmt.Println("It's a weekday")
	}

	hour := time.Now().Hour()
	switch {
	case hour < 6 || hour > 22:
		fmt.Println("Night")
	case hour >= 6 && hour < 12:
		fmt.Println("Before noon")
	case hour >= 12 && hour < 18:
		fmt.Println("After noon")
	default:
		fmt.Println("Evening")
	}

}
