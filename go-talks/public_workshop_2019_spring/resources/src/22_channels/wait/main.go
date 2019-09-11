package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Hello gouroutines!")
		ch <- "done"
	}()

	<-ch
}
