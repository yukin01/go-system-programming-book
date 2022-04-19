package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./q1/out.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "%d\n", 123)
	fmt.Fprintf(file, "%s\n", "string")
	// var num float64 = 1 / 2
	num := float64(1) / float64(2)
	fmt.Fprintf(file, "%f\n", num)
}
