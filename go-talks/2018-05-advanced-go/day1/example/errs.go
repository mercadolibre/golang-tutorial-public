package main

import (
	"errors"
	"fmt"
)

// START 1 OMIT
var (
	ErrFileNotFound = errors.New("file not found")
	ErrNetwork      = errors.New("network error")
)

// END 1 OMIT

// START 2 OMIT
func SomeFunction(argument string) error {
	//return ErrFileNotFound
	return FileNotFound{Filename: argument}
}

// END 2 OMIT

// START 4 OMIT
type FileNotFound struct {
	Filename string
}

// END 4 OMIT

// START 5 OMIT
func (f FileNotFound) Error() string {
	return fmt.Sprintf("filename: %s", f.Filename)
}

// END 5 OMIT

func main() {

	// START 3 OMIT
	err := SomeFunction("argument")
	switch err {
	case ErrFileNotFound:
		//...
	case ErrNetwork:
		//...
	}
	fmt.Printf("%s\n", err)
	// END 3 OMIT

}

func main2() {
	// START 6 OMIT
	if err := SomeFunction("argument"); err != nil {
		if ferr, ok := err.(FileNotFound); ok {
			f.Printf("%s\n", ferr.Filename)
		}
		fmt.Printf("%s\n", err)
	}

	// END 6 OMIT
}

// START 7 OMIT
func DoSomething() {
	defer func() {
		if r := recover(); r != nil { // HLxxx
			fmt.Printf("recovered: %v", r)
		}
	}()
	otherFunction(42)
}

func otherFunction(value int) {
	if value == 42 {
		panic("todo mal.")
	}
}

// END 7 OMIT

func main3() {

	DoSomething()
}
