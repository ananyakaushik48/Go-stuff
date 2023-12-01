package main

import (
	"fmt"
	"io"
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
	// LEGACY WAY TO DO THE SAME THING AS COPY
	// Creating buffers to hold input/output
	// input := make([]byte, 4096)

	// s, err := reader.Read(input)
	// // Use the reader to read bytes into "input" from cli
	// if err != nil {
	// 	log.Fatalln("Couldnt read data")
	// }
	// fmt.Printf("Read %d bytes from stdin\n", s)

	// // Using the writer to write data from "input" back to the console
	// s, err = writer.Write(input)
	// if err != nil {
	// 	log.Fatalln("Couldnt write data")
	// }
	// fmt.Printf("Wrote %d bytes to stdout\n", s)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
