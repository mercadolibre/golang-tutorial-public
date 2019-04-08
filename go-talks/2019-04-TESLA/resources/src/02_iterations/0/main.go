package main

import (
	"fmt"
)

func main() {
	var s string

	args:=[]string{"arg1","arg2","arg3","arg4"}

	for i := 0; i < len(args); i++ {
		s += args[i] + " "
	}

	fmt.Println("Result 1: " + s)
}
