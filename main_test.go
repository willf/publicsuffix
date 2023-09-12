package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestNullableString_MarshalJSON(t *testing.T) {
	// Test that a non-empty NullableString is marshaled to JSON correctly
	s := NullableString("test")
	expected := []byte(`"test"`)
	actual, err := json.Marshal(s)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}

	// Test that an empty NullableString is marshaled to JSON as null
	s = NullableString("")
	expected = []byte("null")
	actual, err = json.Marshal(s)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestParse(t *testing.T) {
	// Test that a valid URL is processed correctly
	input := "http://www.example.com"
	expected := Result{
		Input:        input,
		BasicDomain:  "example.com",
		TopLevel:     "com",
		SecondLevel:  "example",
		ThirdLevel:   "www",
		ErrorMessage: NullableString(""),
	}
	actual := Parse(input)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	input = "www.example.com"
	expected = Result{
		Input:        input,
		BasicDomain:  "example.com",
		TopLevel:     "com",
		SecondLevel:  "example",
		ThirdLevel:   "www",
		ErrorMessage: NullableString(""),
	}
	actual = Parse(input)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	input = "www.example.com/test?query=string"
	expected = Result{
		Input:        input,
		BasicDomain:  "example.com",
		TopLevel:     "com",
		SecondLevel:  "example",
		ThirdLevel:   "www",
		ErrorMessage: NullableString(""),
	}
	actual = Parse(input)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test that an invalid URL returns an error
	input = "not a url"
	expected = Result{
		Input:        input,
		BasicDomain:  "",
		TopLevel:     "",
		SecondLevel:  "",
		ThirdLevel:   "",
		ErrorMessage: NullableString("not a url is a suffix"),
	}
	actual = Parse(input)

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestMain(t *testing.T) {
	// Test that the main function runs without errors
	main()
}
