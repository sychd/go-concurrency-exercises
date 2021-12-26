package main

import (
	"fmt"
)

// ByteCounter type
type ByteCounter int

func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	*bc += ByteCounter(len(p))
	return len(p), err
}

func (bc *ByteCounter) String() string {
	return string(*bc)
}

// TODO: Implement Write method for ByteCounter
// to count the number of bytes written.

func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
