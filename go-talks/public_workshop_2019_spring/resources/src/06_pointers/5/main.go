package main

import "fmt"

func main() {
	a := new(int)
	*a = 10
	fmt.Printf("Before method   -> pointer value: %v, pointer address: %v, value: %d \n", a, &a, *a)
	increment(a)
	fmt.Printf("After method    -> pointer value: %v, pointer address: %v, value: %d \n", a, &a, *a)
}

func increment(a *int) {
	fmt.Printf("In inc method a -> pointer value: %v, pointer address: %v, value: %d \n", a, &a, *a)
	b := new(int)
	*b = *a
	*b++
	a = b
	fmt.Printf("In inc method a -> pointer value: %v, pointer address: %v, value: %d \n", a, &a, *a)
	fmt.Printf("In inc method b -> pointer value: %v, pointer address: %v, value: %d \n", b, &b, *b)
}
