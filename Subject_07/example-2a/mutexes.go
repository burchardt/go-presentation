package main

import (
	"fmt"
	"sync"
	"time"
)

type safeCounter struct {
	i int
	sync.RWMutex
}

func (sc *safeCounter) Increment() {
	sc.Lock()
	defer sc.Unlock()
	sc.i++
}

func (sc *safeCounter) Decrement() {
	sc.Lock()
	defer sc.Unlock()
	sc.i--
}

func (sc *safeCounter) GetValue() int {
	sc.RLock()
	defer sc.RUnlock()
	return sc.i
}

func main() {
	sc := new(safeCounter)

	for i := 0; i < 10000; i++ {
		go sc.Increment()
		go sc.Decrement()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(sc.GetValue())
}
