package main

import (
	"fmt"
)

func main() {

	var sChannel chan string = make(chan string)

	// sChannel <- "hello"

	go func() { sChannel <- "hello" }()

	fmt.Println("Channel output: ", <-sChannel)
}
