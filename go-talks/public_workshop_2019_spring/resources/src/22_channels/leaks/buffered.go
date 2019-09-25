package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

func main() {
	time.Sleep(500 * time.Millisecond)

	start := runtime.NumGoroutine()
	fmt.Printf("Start: %d\n", start)

	_ = UnleakedJob()

	time.Sleep(500 * time.Millisecond)

	end := runtime.NumGoroutine()
	fmt.Printf("End: %d\n", end)

	if end > start {
		fmt.Printf("\nError! Start: %d - End: %d\n", start, end)
	} else {
		fmt.Println("\nOk")
	}
}

// START OMIT
func UnleakedJob() error {
	n := 5
	okChan := make(chan bool, n)
	errChan := make(chan error, n)

	for i := 0; i < n; i++ {
		go doJob(i, okChan, errChan)
	}

	for i := 0; i < n; i++ {
		select {
		case <-okChan:
			// do something
		case err := <-errChan:
			return err
		}
	}
	return nil
}

// END OMIT
func doJob(data int, okChan chan bool, errChan chan error) {
	if data == 1 {
		errChan <- errors.New("something went wrong")
		return
	}
	okChan <- true
}
