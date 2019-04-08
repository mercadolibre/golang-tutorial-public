package main

import "fmt"

// START OMIT
func main() {
	panicAndRecover()
	fmt.Println("I need to run the statement at any cost!")
}

func panicAndRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("Unable to connect database!")
}

// END OMIT
