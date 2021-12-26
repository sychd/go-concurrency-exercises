package main

import (
	"context"
	"fmt"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	generator := func(ctx context.Context) <-chan int {
		ch := make(chan int)

		go func() {
			defer close(ch)
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()

		return ch
	}
	ctx, cancel := context.WithCancel(context.Background())

	numbers := generator(ctx)

	counter := 0
	for v := range numbers {
		fmt.Println(v)
		counter++

		if counter == 5 {
			cancel()
		}
	}
	fmt.Println("Done")
}
