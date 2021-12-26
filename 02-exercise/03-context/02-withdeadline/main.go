package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.

	compute := func(ctx context.Context) <-chan data {
		fmt.Println("Entered fn")
		ch := make(chan data)
		go func() {
			fmt.Println("Entered goroutine")
			defer func() {
				fmt.Println("CLOSE")
				close(ch)
			}()
			// Simulate work.

			deadline, ok := ctx.Deadline()
			if ok {
				fmt.Println("IF ok")
				if deadline.Sub(time.Now().Add(50*time.Millisecond)) < 0 {
					fmt.Println("not sufficient time given... terminating")
					return
				}
			}
			fmt.Println("before sleep")
			time.Sleep(50 * time.Millisecond)
			fmt.Println("after sleep")

			// Report result.
			select {
			case <-ctx.Done():
				fmt.Println("Done")
				return
			case ch <- data{"123"}:
				fmt.Println("Data sent")
			}

		}()
		fmt.Println("Exit fn")
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	deadline := time.Now().Add(55 * time.Millisecond)
	ctx, close := context.WithDeadline(context.Background(), deadline)
	ch := compute(ctx)
	fmt.Println("Before data get M")
	d, ok := <-ch
	fmt.Printf("After data get M, d =  %v, ok = %v\n", d, ok)
	if ok {
		fmt.Printf("M work complete: %s\n", d)
	}
	close()

}
