package main

import (
	"fmt"
	"plugin"
)

// START 1 OMIT
type Say interface {
	Say() string
}

func main() {
	p, err := plugin.Open("plugin.so") // HLxxx
	if err != nil {
		panic(err)
	}
	symbol, err := p.Lookup("SaySomething") // HLxxx
	if err != nil {
		panic(err)
	}
	say, ok := symbol.(Say) // HLxxx
	if ok {
		fmt.Printf("%s\n", say.Say()) // HLxxx
	}
}

// END 1 OMIT
