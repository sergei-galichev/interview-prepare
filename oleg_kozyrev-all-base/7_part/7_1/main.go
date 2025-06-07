// Напишите функцию, которая мерджит n-е количество каналов в один

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	channels := make([]chan int64, 1000)

	for i := range channels {
		channels[i] = make(chan int64)
	}

	for i := range channels {
		go func(i int) {
			channels[i] <- int64(i)
			close(channels[i])
		}(i)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	for v := range merge(ctx, channels...) {
		println(v)
	}
}

func merge(ctx context.Context, channels ...chan int64) chan int64 {
	res := make(chan int64)

	merger := func(c chan int64) {
		//for v := range c {
		//	res <- v
		//}

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancelled")
				return
			case val, ok := <-c:
				if !ok {
					return
				}

				res <- val
			}
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			merger(ch)
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
