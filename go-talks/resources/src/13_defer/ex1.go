package main

import "fmt"

func main() {
	fmt.Println("Counting")
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
