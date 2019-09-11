package main

import (
	"fmt"
)

type Expectator interface {
	GetDiscount() float32
}

type BaseExpectator struct{}

func (*BaseExpectator) GetDiscount() float32 {
	return 0
}

type Retired struct {
	BaseExpectator
}

func (*Retired) GetDiscount() float32 {
	return 0.5
}

type Guess struct {
	BaseExpectator
}

func (*Guess) GetDiscount() float32 {
	return 1
}

func main() {
	expectators := [...]Expectator{&Retired{}, &BaseExpectator{}, &Guess{}}

	fmt.Println(netIncome(expectators[:], 100))
}

func netIncome(expectators []Expectator, ticketPrice float32) float32 {
	var total float32

	for _, e := range expectators {
		total += ticketPrice - e.GetDiscount()*ticketPrice
	}

	return total
}
