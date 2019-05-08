package main

import (
	"fmt"
)

func main() {

	var sChannel chan string = make(chan string)

	go func() {
		sChannel <- "hello"
	}()

	fmt.Println("Channel output: ", <-sChannel)

}
