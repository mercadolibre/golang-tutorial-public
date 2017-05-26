package main

import (
	"fmt"
	"errors"
	"time"
)

// START OMIT
func main() {
	go panicFunc()
	go panicFunc()
	go panicFunc()
	go panicFunc()

	time.Sleep(2*time.Second)
	fmt.Println("Program ends gracefully")
}

func panicFunc(){
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic recovered. Error:%v\n",err)
		}
	}()

	panic(errors.New("forced panic"))
}
// END OMIT

