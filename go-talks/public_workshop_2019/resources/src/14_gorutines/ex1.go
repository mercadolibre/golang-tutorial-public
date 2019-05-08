package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello gouroutines!")
}

func main() {
	go sayHello()
	// Do other things
}
