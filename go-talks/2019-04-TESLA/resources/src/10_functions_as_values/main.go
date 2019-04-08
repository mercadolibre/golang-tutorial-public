package main

import (
	"fmt"
)

type function func(string)

func main() {
	executeFunction(getFunction(), "test")
}

func getFunction() func(string) {
	f := func(value string) {
		fmt.Printf("Execute function with value: %s \n", value)
	}
	return f
}

func executeFunction(f function, value string) {
	f(value)
}
