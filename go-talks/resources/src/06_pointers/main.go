package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 4
	fmt.Println("X Value:" + strconv.Itoa(x))
	y := new(int)
	fmt.Println("Y Value:" + strconv.Itoa(*y))
	*y = 4
	fmt.Println("Y Value:" + strconv.Itoa(*y))
	if y == &x {
		fmt.Println("X & Y are equals")
	} else {
		fmt.Println("X & Y are not equals")
	}
}
