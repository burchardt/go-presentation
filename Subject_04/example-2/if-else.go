package main

import "fmt"

func main() {
	x := 0
	if x > 0 {
		fmt.Println("x is positive")
	} else if x < 0 {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is zero")
	}

	nicknamesMap := map[string]string{"Joanna": "Asia", "Tomasz": "Tomek", "Bartosz": "Bartek"}
	if nickname, ok := nicknamesMap["Joanna"]; ok {
		fmt.Println("Joanna -> ", nickname)
	} else {
		fmt.Println("Unknown name")
	}
	// fmt.Println(nickname)
}
