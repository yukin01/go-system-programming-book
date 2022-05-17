package main

import (
	"crypto/rand"
	"os"
)

func main() {
	buffer := make([]byte, 1024)

	if _, err := rand.Reader.Read(buffer); err != nil {
		panic(err)
	}

	file, err := os.Create("rand.out")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.Write(buffer); err != nil {
		panic(err)
	}
}
