package main

import (
	"fmt"
)

func main() {

	var sChannel chan string = make(chan string, 2)

	sChannel <- "hello"
	sChannel <- "world"
	sChannel <- "!"

	fmt.Printf("%s %s\n", <-sChannel)
}
