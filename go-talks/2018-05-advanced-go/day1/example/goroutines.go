package main

import "fmt"

func main() {
	// START 1 OMIT
	ch := make(chan int)

	go func(count int) {
		for index := 0; index < count; index++ {
			ch <- index
		}
		close(ch)
	}(1485)

	for item := range ch {
		fmt.Printf("item: %d\n", item)
	}
	// END 1 OMIT
}
