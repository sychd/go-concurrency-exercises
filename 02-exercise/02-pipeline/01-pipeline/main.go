package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, num := range nums {
			ch <- num
		}
		close(ch)
	}()

	return ch
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range ch {
			out <- num * num
		}
		close(out)
	}()

	return out
}

func main() {
	for v := range square(generator(1, 2, 3, 4, 5)) {
		fmt.Println(v)
	}
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

}
