package main

import (
	"fmt"
	"github.com/rodriguezgustavo/go-talks/2016-10-BsAs/resources/src/4_packages/math"
)

func main() {
	var result int
	args:=[]int{1,2,3,4}
	for _, val := range args {
		result = math.Add(result, val)
	}
	fmt.Printf("Suma: %d\n", result)
}
