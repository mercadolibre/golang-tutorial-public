package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	delete(m, 3)

	elem := m[2]
	fmt.Printf("Value: %d \n", elem)

	elem, ok := m[2]
	fmt.Printf("Value: %d - def: %t \n", elem, ok)

	elem, ok = m[3]
	fmt.Printf("Value: %d - def: %t", elem, ok)
}
