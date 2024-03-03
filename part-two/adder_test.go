package main

import "testing"

func TestAdder(t *testing.T) {
	sum := Adder(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
