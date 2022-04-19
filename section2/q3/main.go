package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	buffer := gzip.NewWriter(w)
	writer := io.MultiWriter(buffer, os.Stdout)

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	_ = encoder.Encode(source)
	buffer.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(":8080", nil)
}
