package serializemsgpack_test

import (
	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Name     string
	Lastname string
}

// START 1 OMIT
func testMsgPack() {
	person := Person{
		Name:     "Carlos",
		Lastname: "Petruza",
	}
	content, err := msgpack.Marshal(person) // HLxxx
	if err != nil {
		panic(err)
	}
	_ = content
}

// END 1 OMIT
