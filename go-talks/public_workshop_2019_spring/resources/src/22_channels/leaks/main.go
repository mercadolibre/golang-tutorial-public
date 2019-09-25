package main

import (
	"errors"
	"fmt"
)

func main() {
	err := LeakedJob()
	fmt.Println(err)
}

// START OMIT
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

// END OMIT
