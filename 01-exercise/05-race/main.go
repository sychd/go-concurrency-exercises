package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//TODO: identify the data race
// fix the issue.

func main() {
	start := time.Now()
	ch := make(chan bool)

	t := time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		ch <- true
	})

	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}
}

func main2() {
	start := time.Now()
	wg := sync.WaitGroup{}

	t := time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		wg.Done()
	})

	for time.Since(start) < 5*time.Second {
		wg.Add(1)
		t.Reset(randomDuration())
		wg.Wait()
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

//----------------------------------------------------
// (main goroutine) -> t <- (time.AfterFunc goroutine)
//----------------------------------------------------
// (working condition)
// main goroutine..
// t = time.AfterFunc()  // returns a timer..

// AfterFunc goroutine
// t.Reset()        // timer reset
//----------------------------------------------------
// (race condition- random duration is very small)
// AfterFunc goroutine
// t.Reset() // t = nil

// main goroutine..
// t = time.AfterFunc()
//----------------------------------------------------
