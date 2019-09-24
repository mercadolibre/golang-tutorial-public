package main

import "fmt"

func main() {
	n := 5

	fmt.Println("Counting")
	for i := 1; i <= n; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
