package main

import (
	"fmt"
)

type FullName interface {
	FullName() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) FullName() string {
	return p.Name
}

type Employee struct {
	Person
}

type Manager struct {
	Employee
}

type Cashier struct {
	Employee
}

type Supervisor struct {
	Employee
}

func (s Supervisor) FullName() string {
	return "Supervisor: " + s.Person.FullName()
}

func main() {
	p1 := Supervisor{
		Employee: Employee{
			Person: Person{
				Name: "Eduardo",
				Age:  28,
			},
		},
	}

	PrintName(p1)
}

func PrintName(p FullName) {
	fmt.Println(p.FullName())
}
