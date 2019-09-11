package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2

	for key, value := range m {
		fmt.Printf("Key: %d, Value: %d \n", key, value)
	}

	var m2 map[int]int
	fmt.Println(m2==nil) //panic if m2[0]=1
}