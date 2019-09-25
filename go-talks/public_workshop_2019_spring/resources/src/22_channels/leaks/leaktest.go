package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

// START OMIT
func main() {
	time.Sleep(500 * time.Millisecond)

	start := runtime.NumGoroutine()
	fmt.Printf("Start: %d\n", start)

	_ = LeakedJob()

	time.Sleep(500 * time.Millisecond)

	end := runtime.NumGoroutine()
	fmt.Printf("End: %d\n", end)

	if end > start {
		fmt.Printf("\nError! Start: %d - End: %d\n", start, end)
	} else {
		fmt.Println("\nOk")
	}
}

// END OMIT
func LeakedJob() error {
	okChan := make(chan bool)
	errChan := make(chan error)
	for i := 0; i < 5; i++ {
		go doJob(i, okChan, errChan)
	}
	for i := 0; i < 5; i++ {
		select {
		case <-okChan:
			// do something
		case err := <-errChan:
			return err
		}
	}
	return nil
}

func doJob(data int, okChan chan bool, errChan chan error) {
	if data == 1 {
		errChan <- errors.New("something went wrong")
		return
	}
	okChan <- true
}
