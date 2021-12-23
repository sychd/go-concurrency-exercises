package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	var sum int64
	numCPU := runtime.NumCPU()
	// Utilize all cores on machine
	runtime.GOMAXPROCS(numCPU)

	// Run computation for each part in seperate goroutine.
	var wg sync.WaitGroup
	slice := len(numbers) / numCPU

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go (func(i int) {
			right := (i + 1) * slice
			if i+1 == numCPU {
				right = len(numbers)
			}

			// Add part sum to cummulative sum
			sum += Add(numbers[i*slice : right])
			wg.Done()
		})(i)
	}
	wg.Wait()

	return sum
}
