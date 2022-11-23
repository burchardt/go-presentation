package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	maxSleepTimeout = 100
)

func numbers(wg *sync.WaitGroup, mux *sync.RWMutex) {
	mux.Lock()
	defer mux.Unlock()
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Duration(rand.Intn(maxSleepTimeout)) * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	wg.Done()
}

func characters(wg *sync.WaitGroup, mux *sync.RWMutex) {
	mux.Lock()
	defer mux.Unlock()
	for i := 'a'; i <= 'j'; i++ {
		time.Sleep(time.Duration(rand.Intn(maxSleepTimeout)) * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	fmt.Println()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mux sync.RWMutex

	for i := 0; i < 5; i++ {
		wg.Add(2)
		go numbers(&wg, &mux)
		go characters(&wg, &mux)
	}
	wg.Wait()
}
