package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func Increment(n *int32) {
	atomic.AddInt32(n, 1)
}

func Decrement(n *int32) {
	atomic.AddInt32(n, -1)
}

func main() {
	var value int32
	for i := 0; i < 1000; i++ {
		go Increment(&value)
		go Decrement(&value)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(value)
}
