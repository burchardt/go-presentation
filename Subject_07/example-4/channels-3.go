package main

import "fmt"

func nonBlockingReceiver(msgChan chan string) {
	select {
	case msg := <-msgChan:
		fmt.Println("received message: ", msg)
	default:
		fmt.Println("no message received")
	}
}

func nonBlockingSender(msgChan chan string, message string) {
	select {
	case msgChan <- message:
		fmt.Println("sent message: ", message)
	default:
		fmt.Println("no message sent")
	}
}

func main() {
	msgChan := make(chan string)

	nonBlockingSender(msgChan, "hello")
	nonBlockingReceiver(msgChan)
}
