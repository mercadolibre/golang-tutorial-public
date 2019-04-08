package main

import (
	"fmt"
)

func main() {
	a := [2]string{"1", "2"}
	b := [...]string{"1", "2"}
	c := [2]string{"1","3"}

	fmt.Println(c[0], c[len(a)-1] )
	fmt.Println( a == b, a==c)

	for _, value := range a {
		fmt.Print(value+" ")
	}
}
