package main

import (
	"fmt"
)

// Examples

func ExampleAdd() {
	result, _ := add(a, b)

	fmt.Println(result)

	// Output: 6
}

func ExampleSubtract() {
	result, _ := subtract(a, b)

	fmt.Println(result)

	// Output: 2
}

func ExampleMultiply() {
	result, _ := multiply(a, b)

	fmt.Println(result)

	// Output: 8
}

func ExampleDivide() {
	result, _ := divide(a, b)

	fmt.Println(result)

	// Output: 2
}
