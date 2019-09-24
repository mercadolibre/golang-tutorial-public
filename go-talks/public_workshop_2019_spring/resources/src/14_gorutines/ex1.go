package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go sayHello()

	fmt.Println("Good bye!")

	time.Sleep(2 * time.Second)
}

func sayHello() {
	time.Sleep(time.Second)

	fmt.Println("Hello goroutines!")
}

// END OMIT
