package main

import "fmt"

var i int
var f float64
var b bool
var s string

func init() { fmt.Println("Se ejecuta init") }

func main() { fmt.Printf("%v %v %v %q\n", i, f, b, s) }
