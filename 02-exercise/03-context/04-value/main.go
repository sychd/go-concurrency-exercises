package main

import (
	"context"
	"fmt"
)

type database map[string]bool

var db database = database{
	"jane": true,
}

func main() {
	processRequest("jane")
}

func processRequest(userid string) {
	// TODO: send userID information to checkMemberShip through context for
	// map lookup.
	rootCtx := context.WithValue(context.Background(), "userid", userid)
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	ch := checkMemberShip(ctx)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMemberShip - takes context as input.
// extracts the user id information from context.
// spins a goroutine to do map lookup
// sends the result on the returned channel.
func checkMemberShip(ctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		userid := ctx.Value("userid").(string)
		// do some database lookup
		status := db[userid]
		ch <- status
	}()
	return ch
}
