package main

import (
	"fmt"
)

func main() {
	a := 10

	fmt.Printf("Before method    -> content: %d \n", a)
	increment(a)
	fmt.Printf("After  method    -> content: %d \n", a)
}

func increment(a int) {
	a++
	fmt.Printf("Increment method -> content: %d \n", a)
}
