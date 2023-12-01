package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct {
}

// This is the standard pattern to take input
func (foo *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (f *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	// Instantiating reader and writer
	var (
		reader FooReader
		writer FooWriter
	)

	// Creating buffers to hold input/output

	input := make([]byte, 4096)

	s, err := reader.Read(input)
	// User the reader to read input from cli
	if err != nil {
		log.Fatalln("Couldnt read data")
	}

}
