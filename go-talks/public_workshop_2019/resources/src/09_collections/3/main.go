package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int) 
	var m2 map[int]int 

	m[1] = 1
	m[2] = 2
	m[3] = 3
	delete(m,3)

	fmt.Println(m2==nil) //panic if m2[0]=1
	for key, value := range m {
		fmt.Printf("Key: %d, Value: %d \n", key, value)
	}
}
