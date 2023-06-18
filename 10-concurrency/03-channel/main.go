package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Printf("Sent %d to the channel\n", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			val := <-ch
			fmt.Printf("Received %d from the channel\n", val)
		}
	}()

	wg.Wait() // Wait for all goroutines to finish
}
