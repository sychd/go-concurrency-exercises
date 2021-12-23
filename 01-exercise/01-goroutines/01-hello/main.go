package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")
	c := make(chan string)

	// goroutine
	go fun("GR-1")

	// goroutine with anonymous function
	go (func(c chan string) {
		fun("GR-2")
		c <- "done"
	})(c)

	// goroutine with function value call
	funcValue := fun
	go funcValue("GR3")

	// wait for goroutines to end
	<-c

	fmt.Println("done..")
}
