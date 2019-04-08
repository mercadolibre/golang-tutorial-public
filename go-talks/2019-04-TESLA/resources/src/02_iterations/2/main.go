package main

import (
	"fmt"
)

func main() {
	var s string
	args:=[]string{"arg1","arg2","arg3","arg4"}

	i := 0
	for {
		if i >= len(args) {
			break
		}
		s += args[i] + " "
		i++
	}
	fmt.Println("Result 3: " + s)
}
