package main

import (
	"fmt"
	"github.com/mercadolibre/golang-tutorial/go-talks/resources/src/04_packages/math"
)

func main() {
	var result int
	args := []int{1, 2, 3, 4}
	for _, val := range args {
		result = math.Add(result, val)
	}
	fmt.Printf("Suma: %d \n", result)
}
