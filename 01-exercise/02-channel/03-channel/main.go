package main

import (
	"fmt"
)

func main() {
	size := 6
	ch := make(chan int, size)

	go func() {
		defer close(ch)

		// TODO: send all iterator values on channel without blocking
		for i := 0; i < size; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}
