package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var stringNumber = "42"

	a, _ := strconv.ParseInt(stringNumber, 10, 64)
	fmt.Println("String to int: ", a)

	b := int64(3)
	var c = math.Sqrt(float64(a*a + b*b))
	fmt.Println("Float: ", c)

	var z = uint(c)
	fmt.Println("Uint", z)

	var toString = strconv.Itoa(32)
	fmt.Println("Int to string: ", toString)

}
