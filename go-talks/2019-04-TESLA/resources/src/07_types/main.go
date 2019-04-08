package main

import (
	"fmt"
)


// START OMIT

type Celsius float64
type Fahrenheit float64

const AbsoluteZeroC Celsius = -273.15
const FreezingC     Celsius = 0
const BoilingC      Celsius = 100

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// END OMIT

func main() {
	args:= []float64{15,2, 20, 100, 120}
	for _, t := range args {
		f := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, FToC(f).String(), c, CToF(c).String())
	}
}
