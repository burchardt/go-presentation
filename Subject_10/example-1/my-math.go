package my_math

import "errors"

func Plus(a int, b int) int {
	return a + b
}

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func SumAndAverage(nums ...int) (int, float64) {
	total := Sum(nums...)
	average := float64(total) / float64(len(nums))
	return total, average
}

func Divide(nominator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("Cannot Divide by zero")
	}
	return nominator / denominator, nil
}
