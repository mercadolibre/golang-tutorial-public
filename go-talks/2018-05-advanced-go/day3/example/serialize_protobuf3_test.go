package serializegob_test

// START 1 OMIT
import (
	"path/to/example" // HLxxx

	"github.com/golang/protobuf/proto"
)

func testProtobuf3() {
	person := &example.Person{ // HLxxx
		Name:     "Carlos",
		Lastname: "Petruza",
	}
	data, err := proto.Marshal(person) // HLxxx
	if err != nil {
		panic(err)
	}
}

// END 1 OMIT
