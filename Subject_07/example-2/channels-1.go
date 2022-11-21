package main

import (
	"fmt"
	"time"
)

func printer(msg chan string) {
	for {
		message := <-msg
		fmt.Println(message)
	}
}

func main() {
	msgChan := make(chan string)
	go printer(msgChan)
	msgChan <- "Hello"
	time.Sleep(1 * time.Second)
	msgChan <- "Hello Again"
	time.Sleep(1 * time.Second)
}
