package main

import (
	"fmt"
)

type LoggerLevel string

const (
	info    LoggerLevel = "INFO"
	warning LoggerLevel = "WARNING"
	error   LoggerLevel = "ERROR"
)

func getLogger(level LoggerLevel) func(string) {
	return func(message string) {
		fmt.Printf("%s: %s\n", level, message)
	}
}

func doSometing(logger func(string)) {
	logger("Start runner")
	// Do something
	logger("Stop runner")
}

func main() {
	func(text string) {
		fmt.Println(text)
	}("my anonymus function called")

	localFunction := func(text string) {
		fmt.Println(text)
	}
	localFunction("my local function called")
	localFunction("my local function called again")

	logger := getLogger(info)
	doSometing(logger)

	logger("End of main function")
}
