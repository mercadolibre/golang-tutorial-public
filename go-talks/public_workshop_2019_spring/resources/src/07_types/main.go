package main

import (
	"fmt"
)

// START OMIT

type Celsius float64
type Fahrenheit float64

const AbsoluteZeroC Celsius = -273.15
const FreezingC Celsius = 0
const BoilingC Celsius = 100

func (c Celsius) String() string {
	return fmt.Sprintf("%.3g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.3g°F", f)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// END OMIT

func main() {
	args := []float64{15, 2, 20, 100, 120}

	fmt.Printf("f = FToC(f) \t c = CToF(c) \n")

	for _, t := range args {
		f := Fahrenheit(t)
		c := Celsius(t)
		fmt.Printf("%s = %s \t %s = %s \n",
			f, FToC(f).String(), c, CToF(c).String())
	}
}
