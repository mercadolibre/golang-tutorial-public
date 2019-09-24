package main

import "fmt"

func main() {
	n := 5
	fmt.Println("Counting")
	i := 1
	for i <= n {
		fmt.Printf("In function: i=%d\n", i)
		defer func() {
			fmt.Printf("In defer: i=%d\n", i)
			i++
		}()
	}
	fmt.Println("Done")
}
