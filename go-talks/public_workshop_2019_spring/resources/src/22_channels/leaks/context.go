package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"
)

func main() {
	time.Sleep(500 * time.Millisecond)

	start := runtime.NumGoroutine()
	fmt.Printf("Start: %d \n", start)

	_ = UnleakedJob()

	time.Sleep(500 * time.Millisecond)

	end := runtime.NumGoroutine()
	fmt.Printf("End: %d \n", end)

	if end > start {
		fmt.Printf(" \nError! Start: %d - End: %d \n", start, end)
	} else {
		fmt.Println(" \nOk")
	}
}

// START OMIT
func UnleakedJob() error {
	n := 5
	okChan := make(chan bool, n)
	errChan := make(chan error, n)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < n; i++ {
		go doJobCtx(ctx, i, okChan, errChan)
	}

	for i := 0; i < n; i++ {
		select {
		case <-okChan:
			// do something
		case err := <-errChan:
			cancel()
			return err
		case <-time.After(1 * time.Second):
			cancel()
			return errors.New("timeout")
		}
	}
	return nil
}

// END OMIT
func doJobCtx(ctx context.Context, data int, okChan chan bool, errChan chan error) {
	select {
	case <-ctx.Done():
		return
	default:
		if data == 1 {
			errChan <- errors.New("something went wrong")
		}
		okChan <- true
	}
}
