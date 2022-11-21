package main

import (
	"errors"
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func sum(nums ...int) int64 {
	var total int64 = 0
	for _, n := range nums {
		total += int64(n)
	}
	return total
}

func sumAndAverage(nums ...int) (int64, float32) {
	total := sum(nums...)
	average := float32(total) / float32(len(nums))
	return total, average
}

func divide(nominator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return nominator / denominator, nil
}

func minimum(x, y int) (min int) {
	if x < y {
		min = x
	} else {
		min = y
	}
	// Values of min is returned
	return
}

func main() {
	p := plus(2, 3)
	fmt.Println("p:", p)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("s:", s)

	t, a := sumAndAverage(10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Printf("total:%d, average:%.2f\n", t, a)

	if d, err := divide(8, 4); err == nil {
		fmt.Println("d:", d)
	} else {
		fmt.Println(err)
	}

	fmt.Println("minimum: ", minimum(5, 7))
}
