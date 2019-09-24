package main

import (
	"fmt"
)

type ApiError struct {
	Message    string
	StatusCode int
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s - Status Code: %d", e.Message, e.StatusCode)
}

func main() {
	var err error = &ApiError{Message: "Bad Request", StatusCode: 400}
	fmt.Printf("Error: %s \n", err.Error())

	//fmt.Printf("StatusCode: %d", err.StatusCode)
}
