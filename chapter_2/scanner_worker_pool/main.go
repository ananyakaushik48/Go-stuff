package main

import (
	"fmt"
	"sync"
)

// This is a single worker function that takes an integer channel and a waitgroup as params
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	// Creating an integer channel of buffersize 100
	ports := make(chan int, 100)
	// Setting up the waitgroup to sync the channel
	var wg sync.WaitGroup
	// this is the worker processing the ports to print the numbers
	// the numbers are from the channel
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	// this loop is adding/writing the numbers to the port channel
	for i := 0; i < cap(ports); i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
