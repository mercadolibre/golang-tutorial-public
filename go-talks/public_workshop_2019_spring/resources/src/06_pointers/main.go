package main

import "fmt"

func main() {
	var p *int	// Explícita
	fmt.Println(p == nil)

	y := new(int)	// Explícita
	fmt.Printf("Y:%p\n", y)
	fmt.Printf("Y Value: %d\n", *y)

	*y = 4
	fmt.Printf("Y Value: %d\n", *y)

	x := 4		// Implícita
	fmt.Printf("X Value: %d\n", x)

	if y == &x {
		fmt.Println("X & Y are equals")
	} else {
		fmt.Println("X & Y are not equals")
	}
}
