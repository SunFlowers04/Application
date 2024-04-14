package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)

	go worker(ctx, &wg)

	time.Sleep(3 * time.Second)
	cancel()

	wg.Wait()

}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {

		select {
		case <-ctx.Done():
			fmt.Println("context is closed")

			return
		default:
			fmt.Println("working")
		}
		time.Sleep(5 * time.Second)

	}

}
