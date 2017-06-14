package main

import "fmt"

func multiples(p int) (int, int) {
	count, sum := 0, 0

	for i := 0; i <= p; i++ {
		if i%3 == 0 || i%5 == 0 {
			count++
			sum += i
		}
	}

	return count, sum
}

func main() {
	fmt.Println(multiples(0))
	fmt.Println(multiples(1))
	fmt.Println(multiples(5))
	fmt.Println(multiples(9))
}
