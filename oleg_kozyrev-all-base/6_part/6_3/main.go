package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan bool, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(ch)

		time.Sleep(3 * time.Second)
		fmt.Println("Go routine finished")
		ch <- false
	}()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Ticker Tik")
			//ch <- true
		case value := <-ch:
			fmt.Printf("Got value %t\n", value)
			wg.Wait()
			return
		}
	}
}
