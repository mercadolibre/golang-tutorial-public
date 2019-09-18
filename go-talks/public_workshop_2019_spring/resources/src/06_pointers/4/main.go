package main

import "fmt"

func main() {
	a := new(int)
	*a = 10

	fmt.Printf("Before -> pointer address: %v, value: %d \n", a, *a)
	increment(a)
	fmt.Printf("After  -> pointer address: %v, value: %d \n", a, *a)
}

func increment(a *int) {
	fmt.Printf("Inc a  -> pointer address: %v, value: %d \n", a, *a)
	b := new(int)
	*b = *a
	*b++
	a = b
	fmt.Printf("Inc a  -> pointer address: %v, value: %d \n", a, *a)
	fmt.Printf("Inc b  -> pointer address: %v, value: %d \n", b, *b)
}
