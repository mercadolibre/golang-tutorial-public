package main

import (
	"fmt"
)

type SuperHero interface {
	Powers()
}

type Superman struct {
	Alias string
	Age   int
}

func (s *Superman) Powers() {
	fmt.Printf("%s powers: %s", s.Alias, "invulnerability, heat vision, flight, speed")
}

type SuperBoy struct {
	Superman
}

func (s *SuperBoy) Powers() {
	fmt.Printf("%s powers: %s", s.Alias, "flight")
}

func main() {
	var superboy *SuperBoy = &SuperBoy{}

	superboy.Alias = "Superboy"
	superboy.Age = 24

	superboy.Powers()

	fmt.Println("")
}
