package main

import "fmt"

type Secret struct {
	secretValue string
}

type DataEntry struct {
	name  string
	value string
}

func TypeSwitch(x interface{}) {
	switch T := x.(type) {
	case Secret:
		fmt.Print("Secret type; ")
		value := x.(Secret).secretValue
		fmt.Println("value: ", value)
	case DataEntry:
		fmt.Println("DataEntry type; ")
		value := x.(DataEntry).value
		fmt.Println("value: ", value)
	default:
		fmt.Printf("Unsupported type: %T", T)
	}
}

func fetchValue(x interface{}) {
	// Let's assume we exect a DataEntry variable
	// Retrieve value in a safe way
	dataEntry, ok := x.(DataEntry)
	if ok {
		fmt.Println("value: ", dataEntry.value)
	}

	// Retrieve value in a risky way
	value := x.(DataEntry).value
	fmt.Println("value: ", value)

	// This will cause PANIC !!!
	// value = something.(Secret).secretValue
}

func main() {
	mySecret := Secret{"my secret"}
	TypeSwitch(mySecret)

	fetchValue(DataEntry{"id", "001"})
}
