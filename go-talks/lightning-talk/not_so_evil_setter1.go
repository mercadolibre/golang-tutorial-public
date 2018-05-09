package main

import (
	"fmt"
)

// START OMIT
//go:noinline
func notSoEvilSetter(slice []int, index, elem int)[]int{
	//Not so evil bit of code
	slice = append(slice, elem)	

	//Set the value at index of the slice to elem
	slice[index] = elem
		
	//Second not so evil bit of code
	slice = slice[:len(slice)-1]
	fmt.Println("Mi slice final es:\t", slice)

	return slice
}

func main() {
	miSlice := make([]int, 1, 2)
	fmt.Println("Mi slice inicial es:\t", miSlice)
	miSlice = notSoEvilSetter(miSlice, 0, 1)
	fmt.Println("Mi slice final es:\t", miSlice)
}
// END OMIT