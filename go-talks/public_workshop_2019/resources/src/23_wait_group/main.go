package main

import (
	"fmt"
	"sync"
)

func main() {

	// Create a wait group.
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("Hello gouroutines!")
		wg.Done()
	}()

	// Wait until everyone finishes.
	wg.Wait()
}
