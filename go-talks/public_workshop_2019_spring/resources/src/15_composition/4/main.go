package main

import "fmt"

// START OMIT
type Ball struct {
	Radius   int
	Material string
}

type Football struct {
	Ball
	Radius int
}

func (b Ball) Bounce() {
	fmt.Printf("Radius = %d \n", b.Radius)
}

func main() {
	fb := Football{Ball{Radius: 5, Material: "leather"}, 7}
	fb.Bounce()
	fb.Ball.Bounce()
	fmt.Printf("Ball Radius = %d \n", fb.Ball.Radius)
	fmt.Printf("Football Radius = %d \n", fb.Radius)
}

// END OMIT
