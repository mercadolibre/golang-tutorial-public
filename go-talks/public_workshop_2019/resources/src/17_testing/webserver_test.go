package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Tests

func TestWebserverGET(t *testing.T) {
	//create the dummy request
	request, _ := http.NewRequest("GET", "localhost:8080/pipi", nil)
	response := httptest.NewRecorder()

	//execute the request
	processCall(response, request)

	//Check the status code
	checkStatus(t, response, 200)

	//Check the content
	checkContent(t, response, "Test webserver")

	//Check the Header
	checkHeader(t, response, "Go-Test", "Test 1.0")

	println("TestWebserverGET OK")
}

func TestWebserverPOST(t *testing.T) {
	//create the dummy request
	request, _ := http.NewRequest("POST", "localhost:8080/pipi", nil)
	response := httptest.NewRecorder()

	//execute the request
	processCall(response, request)

	//Check the status code
	checkStatus(t, response, 401)

	//Check the Header
	checkHeader(t, response, "Unauthorized", "You need to have a valid token")

	println("TestWebserverPOST OK")
}
