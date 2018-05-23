package serializejson_test

import (
	"encoding/json"
	"testing"
)

// START 1 OMIT
type Person struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Level    int    `json:"-"`
}

var data = []byte(`{"name": "carlos","lastname": "petruza"}`)

func TestJson(t *testing.T) {
	var person Person
	// Transform JSON (as bytes) to a Go struct
	if err := json.Unmarshal(data, &person); err != nil { // HLxxx
		panic(err)
	}
	// Transform Go struct to JSON (as bytes)
	bytes, err := json.Marshal(person) // HLxxx
}

// END 1 OMIT
