package utils

import (
	"net/url"
	"testing"
)

func TestEncode(t *testing.T) {
	t.Parallel()
	v := url.Values{}
	v.Add("key1", "value1")
	v.Add("key2", "value2")
	v.Add("key3", "value3")
	keys := []string{"key1", "key2"}
	expected := "key1=value1&key2=value2"
	result := Encode(v, keys)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestConstructFqdnWithAtSymbolAndName(t *testing.T) {

	expected := "example.com."
	result := ConstructFqdn("@", "example.com")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestConstructFqdnWithDotSymbolAndName(t *testing.T) {

	expected := "example.com."
	result := ConstructFqdn(".", "example.com")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestConstructFqdnWithAtSymbolAndNameAndTrailingDot(t *testing.T) {

	expected := "example.com."
	result := ConstructFqdn("@", "example.com.")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestConstructFqdnWithSubDomain(t *testing.T) {

	expected := "www.example.com."
	result := ConstructFqdn("www", "example.com")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestConstructFqdnWithFdqn(t *testing.T) {

	expected := "example.com."
	result := ConstructFqdn("example.com.", "example.com")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestConstructFqdnWithFdqnAndSubDomain(t *testing.T) {

	expected := "www.example.com."
	result := ConstructFqdn("www.example.com.", "example.com")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestConstructFqdnWithNotFdqnAndSubDomain(t *testing.T) {

	expected := "www.example.com.example.com."
	result := ConstructFqdn("www.example.com", "example.com")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestDeconstructFqdnApex(t *testing.T) {
	expected := "@"
	result := DeconstructFqdn("example.com.", "example.com.")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestDeconstructFqdnSubDomain(t *testing.T) {
	expected := "www"
	result := DeconstructFqdn("www.example.com", "example.com")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
