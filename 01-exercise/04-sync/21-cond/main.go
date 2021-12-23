package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	mu.Lock()
	go func() {
		for len(sharedRsc) == 0 {
			fmt.Println("hey")
			c.Wait() // if you comment wait out, you can see "hey" spam
		}

		fmt.Println(sharedRsc["rsc1"])
		wg.Done()
	}()
	time.Sleep(100 * time.Millisecond)
	sharedRsc["rsc1"] = "foo"
	c.Signal()
	wg.Wait()
}
