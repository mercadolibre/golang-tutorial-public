package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("world")

	fmt.Println("hello")

	time.Sleep(time.Second)
}
