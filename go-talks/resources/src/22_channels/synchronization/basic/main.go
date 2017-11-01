package main

import (
	"fmt"
	"time"
)

// read only channel <-chan bool / write only channel chan<- bool
func worker(done chan bool) {
	fmt.Print("working...")

	time.Sleep(time.Second)

	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)

	<-done
}
