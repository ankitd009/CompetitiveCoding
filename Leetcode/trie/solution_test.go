package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	obj := Constructor()
	obj.Insert("hello")
	op := obj.Search("hello")
	if !op {
		t.Errorf("Expected: %v, Got: %v", true, op)
	}
}

func TestStartsWith(t *testing.T) {
	obj := Constructor()
	obj.Insert("hello")
	op := obj.StartsWith("hel")
	if !op {
		t.Errorf("Expected: %v, Got: %v", true, op)
	}
}
