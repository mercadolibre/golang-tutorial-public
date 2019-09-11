package main

import (
	"fmt"
)

func main() {
	//run escape analysis: go run -gcflags '-m' main.go 
	sum := sum()
	fmt.Printf("main - sum: %d, address: %d \n", sum, &sum)
}

func sum() (int){ 
	numbers := make([]int, 3)
	for i:= range numbers { 
		numbers[i] = i + 1
	}
	var sum int
	for _, i := range numbers {
		sum += i
	}
	fmt.Println(numbers)
	return sum
}

func dosomething(numbers [] int){
	for i:= range numbers { 
		numbers[i] = i + 1
	}
}
