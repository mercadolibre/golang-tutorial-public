package main

import (
	"fmt"
)

func main() {
	a := new(int)

	*a = 10

	fmt.Printf("Before method -> address: %v, pointer address: %v, value: %d \n", a, &a, *a)

	increment(a)

	fmt.Printf("Before method -> address: %v, pointer address: %v, value: %d \n", a, &a, *a)
}

func increment(a *int) {
	*a++

	fmt.Printf("In increment method -> address: %v, pointer address: %v, value: %d \n", a, &a, *a)
}
