package main

import "fmt"

type Ball struct {
	Radius   int
	Material string
}

func (b Ball) Bounce() {
	fmt.Printf("Bouncing ball %+v\n", b)
}

type Football struct {
	Ball
}

func main() {
	fb := Football{Ball{Radius: 5, Material: "leather"}}
	fb.Bounce()
}
