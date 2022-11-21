package main

import (
	"fmt"
	"time"
)

func startPrinter(msg chan string) func() {
	stop := make(chan bool)

	go func() {
		fmt.Println("*** Printer started ***")
		for {
			select {
			case message := <-msg:
				fmt.Println(message)
			case <-stop:
				return
			}
		}
	}()

	return func() {
		fmt.Println("*** Printer stopped ***")
		stop <- true
	}
}

func main() {
	msgChan := make(chan string)
	stopPrinter := startPrinter(msgChan)
	msgChan <- "Hello"
	time.Sleep(1 * time.Second)
	msgChan <- "Hello Again"
	time.Sleep(1 * time.Second)
	stopPrinter()
}
