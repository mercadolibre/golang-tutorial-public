package serializegob_test

import (
	"bytes"
	"encoding/gob"
)

type Person struct {
	Name     string
	Lastname string
}

// START 1 OMIT
func testGob() {
	var wire bytes.Buffer
	encoder := gob.NewEncoder(&wire) // HLxxx
	person := Person{
		Name:     "Carlos",
		Lastname: "Petruza",
	}
	if err := encoder.Encode(person); err != nil { // HLxxx
		panic(err)
	}
}

// END 1 OMIT
