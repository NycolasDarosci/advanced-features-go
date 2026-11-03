package main

import (
	"context" //immutable
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("firing 3 requests..")

	// unsynchronized routines
	//go makeRequest(1)
	//go makeRequest(2)
	//go makeRequest(3)

	// wait for all requests to complete
	// time.Sleep(5 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			makeRequest(ctx, id)
		}(i)
	}

	// we can have a request which it will gonna take 10 seconds
	// wg.Wait() will going to wait for all requests to complete
	// timeouts 10 sec -> request timed out
	wg.Wait()

	fmt.Println("Done!")
}

func makeRequest(ctx context.Context, id int) {
	fmt.Printf("[req %d] making request..\n", id)

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Printf("[req %d] request done!\n", id)
		break
	case <-ctx.Done():
		fmt.Println("Requests has timed out!")
	}
}
