package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("./q2/out.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file) // bufio.Writer interface

	_ = writer.Write([]string{"1", "2", "3"})
	_ = writer.Write([]string{"a", "b", "c"})
	writer.Flush()
}
