package main

import (
	"fmt"
	"time"
)

func sayHello2() {
	fmt.Println("Hello goroutines!")
}

func main() {
	go sayHello2()
	// Do other things
	time.Sleep(time.Second)
}
