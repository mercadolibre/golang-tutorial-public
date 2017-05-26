package main

import (
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var a int = 4
var b int = 2

// Execute before and after suit test
func TestMain(m *testing.M) {
	go createWebserver()

	sleep()

	code := m.Run()

	os.Exit(code)
}

//check the body content of a response
func checkContent(t *testing.T, response *httptest.ResponseRecorder, content string) {

	body, _ := ioutil.ReadAll(response.Body)
	if string(body) != content {
		t.Fatalf("Non-expected content %s, expected %s", string(body), content)
	}
}

//Check the status code
func checkStatus(t *testing.T, response *httptest.ResponseRecorder, expected int) {
	if response.Code != expected {
		t.Fatalf("Non-expected status code %v :\n\tbody: %v", expected, response.Code)
	}
}

//Check the specific header value
func checkHeader(t *testing.T, response *httptest.ResponseRecorder, header string, value string) {
	if response.Header().Get(header) != value {
		t.Fatalf("Header: %s, get:%s expected:%s", header, response.Header().Get(header), value)
	}
}

func sleep() {
	time.Sleep(500 * time.Millisecond)
}
