package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// START OMIT
type Shape interface {
	Area() float64
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %f \n", s.Area())
}

func main() {
	r := Rectangle{Width: 12, Height: 6}
	c := Circle{Radius: 10}
	PrintArea(r)
	PrintArea(c)
}

// END OMIT
