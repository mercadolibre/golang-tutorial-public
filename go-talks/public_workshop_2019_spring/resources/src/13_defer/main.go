package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer panic("pannnnic")
	defer fmt.Println("world3")
	fmt.Println("hello")
}
