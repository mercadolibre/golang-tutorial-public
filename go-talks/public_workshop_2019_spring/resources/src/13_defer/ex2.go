package main

import "fmt"

func main() {
	reverseCount(5)
}

func reverseCount(n int) {
	fmt.Println("Counting")
	for i := 1; i <= n; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
