package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxSleepTimeout = 500
)

func numbers(wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Duration(rand.Intn(maxSleepTimeout)) * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func characters(wg *sync.WaitGroup) {
	for i := 'a'; i <= 'j'; i++ {
		time.Sleep(time.Duration(rand.Intn(maxSleepTimeout)) * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go numbers(&wg)
	go characters(&wg)
	wg.Wait()

	fmt.Println()
	fmt.Println("main terminated")
}
