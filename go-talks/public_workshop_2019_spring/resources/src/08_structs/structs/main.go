package main

import "fmt"


type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X:1, Y:2}  // Vertex{1, 2}
	v3 = Vertex{Y:2, X:1}  // Vertex{1, 2}
	v4 = Vertex{X: 1}  // Y:0 is implicit - Vertex{1, 0}
	v5 = Vertex{Y: 2}  // Y:0 is implicit - Vertex{0, 2}
	v6 = Vertex{}      // X:0 and Y:0 - Vertex{0, 0}
	p  = &Vertex{1, 2} // has type *Vertex
)


func main() {
	fmt.Println(v1.X)
	v1.X = 4
	fmt.Println(v1.X)
	fmt.Println(p.X)
}

