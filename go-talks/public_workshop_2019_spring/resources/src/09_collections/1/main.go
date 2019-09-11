package main

import "fmt"

func main() {
	var d [3]int
	fmt.Println(d)

	a := [2]string{"1", "2"}
	b := [...]string{"1", "2"}
	c := [2]string{"1","3"}

	fmt.Println( a == b, a==c)
	fmt.Println(len(a))
	fmt.Println(c[0], c[len(c)-1] )

	for _, value := range a {
		fmt.Print(value+" ")
	}
}
