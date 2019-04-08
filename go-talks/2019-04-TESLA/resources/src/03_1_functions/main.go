package main

import "fmt"

func concat(x, y string) (concat string) {
	concat = x + " " + y
	return
}

func main() {
	a := concat("hello", "world")
	fmt.Println(a)
}
