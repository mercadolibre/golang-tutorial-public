package main

import (
	"fmt"
)

// START OMIT
//go:noinline
func evilSetter(slice []int, index, elem int){
	//Evil bit of code
	slice = append(slice, elem)	

	//Set the value at index of the slice to elem
	slice[index] = elem
		
	//Second evil bit of code
	slice = slice[:len(slice)-1]
	fmt.Println("Mi slice intermedio es:\t", slice)
}

func main() {
	//miSlice := make([]int, 1, 2)
	miSlice := make([]int, 1)
	fmt.Println("Mi slice inicial es:\t", miSlice)
	evilSetter(miSlice, 0, 1)
	fmt.Println("Mi slice final es:\t", miSlice)
}
// END OMIT