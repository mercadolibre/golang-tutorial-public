package main

import (
	"testing"
)

// Tests

func TestAdd(t *testing.T) {
	//t.Parallel()

	result, err := add(a, b)

	if err != nil {
		t.Fatal("We got some error", err)
	}

	expected := 6

	if result != expected {
		t.Fatalf("%d + %d Should be equals to %d", a, b, expected)
	}

	sleep()

	println("TestAdd OK")
}

func TestSubtract(t *testing.T) {
	//t.Parallel()

	result, err := subtract(a, b)

	if err != nil {
		t.Fatal("We got some error", err)
	}

	expected := 2

	if result != expected {
		t.Fatalf("%d - %d Should be equals to %d", a, b, expected)
	}

	sleep()

	println("TestSubtract OK")
}

func TestMultiply(t *testing.T) {
	//t.Parallel()

	result, err := multiply(a, b)

	if err != nil {
		t.Fatal("We got some error", err)
	}

	expected := 8

	if result != expected {
		t.Fatalf("%d * %d Should be equals to %d", a, b, expected)
	}

	sleep()

	println("TestMultiply OK")
}

func TestDivide(t *testing.T) {
	//t.Parallel()

	result, err := divide(a, b)

	if err != nil {
		t.Fatal("We got some error", err)
	}

	expected := 2

	if result != expected {
		t.Fatalf("%d / %d Should be equals to %d", a, b, expected)
	}

	sleep()

	println("TestDivide OK")
}

func TestOperationBigNumberShouldFail(t *testing.T) {
	_, err := multiply(11, 3)

	if err == nil {
		t.Fatal("Invalid big number")
	}

	println("TestOperationBigNumberShouldFail OK")
}

func TestDivideByZeroShouldFail(t *testing.T) {
	_, err := divide(a, 0)

	if err == nil {
		t.Fatal(err.Error())
	}

	println("TestDivideByZeroShouldFail OK")
}
