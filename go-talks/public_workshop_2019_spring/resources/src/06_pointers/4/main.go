package main

import "fmt"

func main() {
	a := new(int)
	*a = 10

	fmt.Printf("Before -> ptr value: %v, content: %d \n", a, *a)
	increment(a)
	fmt.Printf("After  -> ptr value: %v, content: %d \n", a, *a)
}

func increment(a *int) {
	fmt.Printf("Inc a  -> ptr value: %v, content: %d \n", a, *a)
	b := new(int)
	*b = *a
	*b++
	a = b
	fmt.Printf("Inc b  -> ptr value: %v, content: %d \n", b, *b)
	fmt.Printf("Inc a  -> ptr value: %v, content: %d \n", a, *a)
}
