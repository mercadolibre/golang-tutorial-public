package main

import (
	"fmt"
)

func main() {
	a := new(int)
	*a = 10

	fmt.Printf("Before method -> address: %v, value: %d \n", a, *a)
	increment(a)
	fmt.Printf("Before method -> address: %v, value: %d \n", a, *a)
}

func increment(a *int) {
	fmt.Printf("In increment method a -> address: %v, value: %d \n", a, *a)
	b := new(int)
	*b = *a
	*b++
	a = b
	fmt.Printf("In increment method a -> address: %v, value: %d \n", a, *a)
	fmt.Printf("In increment method b -> address: %v, value: %d \n", b, *b)
}
