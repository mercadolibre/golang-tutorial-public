package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	// Create a wait group.
	var wg sync.WaitGroup

	wg.Add(3)
	for i := 0; i < 3; i++ {

		go func() {
			defer wg.Done()
			// Simple function call.
			doJob()
		}()
	}
	// Wait until everyone finishes.
	wg.Wait()
}

func doJob() {
	fmt.Println("Hello goroutines!")
}

//END OMIT
