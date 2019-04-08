package main

import (
	"errors"
	"net/http"
)

/*

	Run All -> go test -bench=.

	Run test only -> go test -run Test*
	Run benchmark only -> go test -run Benchmark* -bench=.
	Run benchmark only parallel -> go test -run Benchmark* -bench=. -cpu 4
	Run examples only -> go test -run Example*

*/

func main() {
	println("Math Methods: ")

	a := 4
	b := 2

	println("A: ", a)
	println("B: ", b)

	result, err := add(a, b)
	if err != nil {
		println(err)
	}

	println("Add: ", result)

	result, err = subtract(a, b)
	if err != nil {
		println(err)
	}

	println("Subtract: ", result)

	result, err = multiply(a, b)
	if err != nil {
		println(err)
	}

	println("Multiply: ", result)

	result, err = divide(a, b)
	if err != nil {
		println(err)
	}

	println("Divide: ", result)

	go createWebserver()
}

func add(a int, b int) (int, error) {
	if a > 10 || b > 10 {
		return 0, errors.New("Numbers too big")
	}

	return a + b, nil
}

func subtract(a int, b int) (int, error) {
	if a > 10 || b > 10 {
		return 0, errors.New("Numbers too big")
	}

	return a - b, nil
}

func multiply(a int, b int) (int, error) {
	if a > 10 || b > 10 {
		return 0, errors.New("Numbers too big")
	}

	return a * b, nil
}

func divide(a int, b int) (int, error) {
	if a > 10 || b > 10 {
		return 0, errors.New("Numbers too big")
	}

	if b == 0 {
		return 0, errors.New("You can not divide by zero")
	}

	return a / b, nil
}

func createWebserver() {
	http.HandleFunc("/", processCall)
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		println("Error!!", err)
	}
}

func processCall(w http.ResponseWriter, r *http.Request) {
	headerMap := w.Header()

	headerMap.Add("Go-Test", "Test 1.0")

	switch r.Method {
	case "GET":
		w.Write([]byte("Test webserver"))

	case "POST":
		headerMap.Add("Unauthorized", "You need to have a valid token")
		w.WriteHeader(401)
	}
}
