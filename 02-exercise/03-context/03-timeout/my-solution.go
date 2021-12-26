package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	// TODO: set a http client timeout
	timeout := 2000 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	wg := sync.WaitGroup{}
	req, err := http.NewRequest("GET", "https://andcloud.io", nil)

	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func() {
		fmt.Println("GR started")
		resp, err := http.DefaultClient.Do(req)
		defer resp.Body.Close()
		defer wg.Done()
		defer fmt.Println("GR Done")

		if err != nil {
			log.Println("ERROR:", err)
			return
		}

		fmt.Println("Before select")
		select {
		case <-ctx.Done():
			fmt.Println("Timeout...")
			return
		default:
			fmt.Println("Fetched!")
			//io.Copy(os.Stdout, resp.Body)
		}
		fmt.Println("After select")
	}()

	wg.Wait()
}
