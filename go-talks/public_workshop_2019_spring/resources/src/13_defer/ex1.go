package main

import "fmt"

// START OMIT
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
// END OMIT