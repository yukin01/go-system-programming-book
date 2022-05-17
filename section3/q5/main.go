package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func copyN(dest io.Writer, src io.Reader, length int) (written int64, err error) {
	lsrc := io.LimitReader(src, int64(length))
	return io.Copy(dest, lsrc)
}

func main() {
	// bytes.Buffer を使って copyN の挙動を確かめる
	source := "TooLongString"
	length := 3
	expected := "Too"

	var dest bytes.Buffer
	n, err := copyN(&dest, bytes.NewBufferString(source), length)
	if err != nil {
		panic(err)
	}

	if n != int64(length) || dest.String() != expected {
		fmt.Printf("copyN() = %v, want %v", dest.String(), expected)
		os.Exit(1)
	}
}
