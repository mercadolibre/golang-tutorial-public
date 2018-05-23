package main

// START 1 OMIT
type SayHello struct{}

func (s *SayHello) Say() string {
	return "hello"
}

var SaySomething SayHello // HLxxx
// END 1 OMIT
