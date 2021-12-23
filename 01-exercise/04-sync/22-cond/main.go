package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		c.Wait()
		fmt.Println(sharedRsc["rsc1"])
		mu.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		c.Wait()
		fmt.Println(sharedRsc["rsc2"])
		mu.Unlock()
	}()

	time.Sleep(500 * time.Millisecond)
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	time.Sleep(500 * time.Millisecond)
	c.Broadcast()

	wg.Wait()
}
