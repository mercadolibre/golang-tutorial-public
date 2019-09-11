package main

import (
	"fmt"
	"os"
	"strconv"
)

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
	if len(os.Args) >= 2 {
		if p, error := strconv.Atoi(os.Args[1]); error == nil {
			count, sum := multiples(p)

			fmt.Printf("Result count: %d, sum: %d \n", count, sum)

		} else {
			fmt.Println("Invalid arguments..., error: ", error.Error())
		}

	} else {
		fmt.Println("Invalid arguments..., must specify p parameter")
	}

}
