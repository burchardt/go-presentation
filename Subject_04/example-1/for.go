package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("i: ", i)
	}

	j := 3
	for j > 0 {
		fmt.Println("j: ", j)
		j -= 1
	}

	x := 0
	for {
		if x == 10 {
			fmt.Println("x: ", x)
			break
		}
		x += 1
	}

	sum := 0
	nums := []int{0, 1, 2, 3, 4}
	for _, n := range nums {
		sum += n
	}
	fmt.Println("sum: ", sum)

	numbers := map[string]int{"x": 101, "y": 102, "z": 103}
	for key, value := range numbers {
		fmt.Printf("numbers[%s]:%d\n", key, value)
	}
}
