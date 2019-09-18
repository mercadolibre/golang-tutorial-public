package main

import "fmt"

func main() {
	a := new(int)
	*a = 10

	fmt.Printf("Before method    -> content: %d \n", *a)
	increment(a)
	fmt.Printf("After  method    -> content: %d \n", *a)
}

func increment(a *int) {
	b := new(int)
	*b = *a
	*b++

	a = b

	fmt.Printf("Increment method -> content: %d \n", *a)
}
