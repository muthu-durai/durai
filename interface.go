package main

import (
	"fmt"
)

func main() {
	var w Writer
	w = outputWriter{}
	w.write([]byte("Hello"))
}

type Writer interface {
	write([]byte) (int, error)
}

type outputWriter struct {
}

func (ow outputWriter) write(data []byte) (int, error) {
	// fmt.Println((data))
	n, err := fmt.Println(string(data))
	return n, err
}