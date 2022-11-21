package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(src, dst string) (int64, error) {
	source, err := os.Open(src)
	defer source.Close()
	if err != nil {
		return 0, err
	}

	destination, err := os.Create(dst)
	defer destination.Close()
	if err != nil {
		return 0, err
	}

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main() {
	if cnt, err := copyFile("source.txt", "destination.txt"); err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println("Bytes copied:", cnt)
	}
}
