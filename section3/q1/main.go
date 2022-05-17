package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("This command requires 2 args.")
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create(flag.Arg(1))
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	if _, err = io.Copy(newFile, file); err != nil {
		panic(err)
	}
}
