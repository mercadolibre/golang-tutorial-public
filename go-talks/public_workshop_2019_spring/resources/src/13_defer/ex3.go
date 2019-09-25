package main

import "fmt"

// START OMIT
func main() {
	leakedCount(2)
}

func leakedCount(n int) {
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
// END OMIT