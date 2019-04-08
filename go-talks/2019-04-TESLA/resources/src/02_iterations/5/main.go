package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now()
	switch {
	case day.Hour() < 12:
		fmt.Println("Good morning!")
	case day.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
