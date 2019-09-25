package main

import (
	"fmt"
)

// START OMIT
func main() {

	var sChannel = make(chan bool)

	go func() {
		result := doJob()

		sChannel <- result
	}()

	fmt.Println("Channel output:", <-sChannel)

}

func doJob() bool {
	// Do some task and return the result
	return true
}

// END OMIT
