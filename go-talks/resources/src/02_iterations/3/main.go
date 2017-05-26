package main

import (
	"fmt"
)

func main() {
	var s string
	args:=[]string{"arg1","arg2","arg3","arg4"}
	
	for _, value := range args {
		s += value + " "
	}
	fmt.Println("Result 4: " + s)
}
