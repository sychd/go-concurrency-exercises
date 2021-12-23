package main

import (
	"fmt"
)

// TODO: Implement relaying of message with Channel Direction

func genMsg(msg string, out chan<- string) {
	// send message on ch1
	out <- msg
}

func relayMsg(in <-chan string, out chan<- string) {
	msg := <-in
	out <- msg
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg
	go genMsg("hi!", ch1)
	go relayMsg(ch1, ch2)

	fmt.Println(<-ch2)
}
