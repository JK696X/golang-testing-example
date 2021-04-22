// +build unit
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//First part Introduction to testing
func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{999, 1001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, recieved: {}",
				test.input, test.expected, output)
		}
	}
}

/* ADVANCED TESTING TECHNIQUE BEGINNING */

/* TEST TABLE EXAMPLE */
type AddResult struct {
	x        int
	y        int
	expected int
}

/* This is an array of test cases we are going to put into TestAdd */
var addResults = []AddResult{
	{1, 1, 2},
}

func TestAdd(t *testing.T) {
	/* For the range of all of our test cases in 'addResults', run them
	through the 'Add' function, get our result, then see if it matches
	the expected result within this element of AddResult */
	for _, test := range addResults {
		result := Add(test.x, test.y)
		if result != test.expected {
			t.Fatal("Expected Result not given")
		}
	}
}

/* Test DIRECTORY EXAMPLE */

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}
	if string(data) != "hello world from test.data" {
		t.Fatal("String contents do not match expected")
	}
}

/* Test HTTP Example */
func TestHTTPRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{ \"status\": \"good\" }")
	}

	r := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Here is our response code: %v\n", string(body))
	if 200 != resp.StatusCode {
		t.Fatal("Status Code not okay")
	}
}

/* ADVANCED TESTING TECHNIQUE END */
