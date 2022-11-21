package main

import "fmt"

func safeReceiver(msgChan chan string) {
	select {
	case msg := <-msgChan:
		fmt.Println("received message: ", msg)
	default:
		fmt.Println("no message received")
	}
}

func safeSender(msgChan chan string, message string) {
	select {
	case msgChan <- message:
		fmt.Println("sent message: ", message)
	default:
		fmt.Println("no message sent")
	}
}

func main() {
	msgChan := make(chan string)

	safeSender(msgChan, "hello")
	safeReceiver(msgChan)
}
