package main

import "fmt"

func main() {
	a := new(int)
	*a = 10
	fmt.Printf("Before -> ptr value: %v, content: %d, ptr address: %v \n", a, *a, &a)
	increment(a)
	fmt.Printf("After  -> ptr value: %v, content: %d, ptr address: %v \n", a, *a, &a)
}

func increment(a *int) {
	fmt.Printf("Inc a  -> ptr value: %v, content: %d, ptr address: %v \n", a, *a, &a)
	b := new(int)
	*b = *a
	*b++
	a = b
	fmt.Printf("Inc b  -> ptr value: %v, content: %d, ptr address: %v \n", b, *b, &b)
	fmt.Printf("Inc a  -> ptr value: %v, content: %d, ptr address: %v \n", a, *a, &a)
}
