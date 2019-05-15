package main

import "fmt"

func main() {

	// Lambda
	func(n string) {
		fmt.Println("Hello", n)
	}("Gophers")

	// Closure
	value := "Gophers"
	func() {
		fmt.Println("Hello", value)
	}()

}
