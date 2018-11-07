package main

import (
	"fmt"
)

// Interface
type Estrategia interface {
    RealizarOperacion(int, int) int
}

// Estrategia Suma
type EstrategiaSuma struct{}

func (e EstrategiaSuma) RealizarOperacion(num1 int, num2 int) int {
    return num1 + num2
}

// Estrategia Resta
type EstrategiaResta struct{}

func (e EstrategiaResta) RealizarOperacion(num1 int, num2 int) int {
    return num1 - num2
}

// Estrategia Multiplica
type EstrategiaMultiplica struct{}

func (e EstrategiaMultiplica) RealizarOperacion(num1 int, num2 int) int {
    return num1 * num2
}

// Contexto
type Contexto struct {
    estrategia Estrategia
}

func (c *Contexto) EjecutarOperacion(num1 int, num2 int) int {
    return c.estrategia.RealizarOperacion(num1, num2)
}

func main() {
	var contexto Contexto
	num1 := 10
	num2 := 5

	contexto = Contexto{EstrategiaSuma{}}
	fmt.Printf("%d + %d = %d\n", num1, num2, contexto.EjecutarOperacion(num1, num2))

	contexto = Contexto{EstrategiaResta{}}
	fmt.Printf("%d - %d = %d\n", num1, num2, contexto.EjecutarOperacion(num1, num2))

	contexto = Contexto{EstrategiaMultiplica{}}
	fmt.Printf("%d * %d = %d\n", num1, num2, contexto.EjecutarOperacion(num1, num2))
}