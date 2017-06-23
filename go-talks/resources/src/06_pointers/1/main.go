package main

import (
	"fmt"
)

func main() {
	a := 10

	fmt.Printf("Before method -> value: %d \n", a)
	increment(a)
	fmt.Printf("After method -> value: %d \n", a)
}

func increment(a int) {
	a++
	fmt.Printf("In increment method -> value: %d \n", a)
}
