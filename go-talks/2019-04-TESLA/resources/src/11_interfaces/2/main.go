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

type Flash struct {
	Alias string
	Age   int
}

func (s *Flash) Powers() {
	fmt.Printf("%s powers: %s", s.Alias, "speed")
}

func main() {
	var superman SuperHero = &Superman{Alias: "Superman", Age: 28}

	superman.Powers()

	fmt.Println("")

	var flash SuperHero = &Flash{Alias: "Flash", Age: 24}

	flash.Powers()

	fmt.Println("")
}
