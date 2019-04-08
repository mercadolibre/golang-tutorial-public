package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("Hello with anonymous functions!")
	}()

	sayHello := func() {
		fmt.Println("Hello with assigned anonymous functions!")
	}

	go sayHello()
	// Do other things
	time.Sleep(time.Second)
}
