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

	fmt.Println("Channel output: ", <-sChannel)

}

func doJob() bool {
	return true
}

// END OMIT
