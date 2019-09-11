package main

import "fmt"

// START OMIT
type fn func(string)

func main() {
	f := getFunction()
	f("test")
	executeFunction(getFunction(), "test")
}

func getFunction() func(string) {
	f := func(value string) {
		fmt.Printf("Execute function with value: %s \n", value)
	}
	return f
}

func executeFunction(f fn, value string) {
	f(value)
}

// END OMIT
