package main

import (
	"fmt"
	"go-concurrency-exercises/01-exercise/01-goroutines/04-add/counting"
	"time"
)

func main() {
	numbers := counting.GenerateNumbers(1e7)

	t := time.Now()
	sum := counting.Add(numbers)
	fmt.Printf("Sequential Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))

	t = time.Now()
	sum = counting.AddConcurrent(numbers)
	fmt.Printf("Concurrent Add, Sum: %d,  Time Taken: %s\n", sum, time.Since(t))
}
