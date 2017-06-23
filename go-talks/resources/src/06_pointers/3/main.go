package main

import (
	"fmt"
)

func main() {
	a := new(int)
	*a = 10

	fmt.Printf("Before method -> value: %d \n", *a)
	increment(a)
	fmt.Printf("After method -> value: %d \n", *a)
}

func increment(a *int) {
	b := new(int)
	*b = *a
	*b++

	a = b

	fmt.Printf("In increment method -> value: %d \n", *a)
}
