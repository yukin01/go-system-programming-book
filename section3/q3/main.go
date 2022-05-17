package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("new.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipwriter := zip.NewWriter(file)
	defer zipwriter.Close()

	writer, err := zipwriter.Create("newFile.txt")
	if err != nil {
		panic(err)
	}

	reader := strings.NewReader("Hello, world!")
	if _, err = io.Copy(writer, reader); err != nil {
		panic(err)
	}
}
