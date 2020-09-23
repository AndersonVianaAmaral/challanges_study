package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	var result = Greeting("BIR")

	if result != "<b>BIR</b>" {
		t.Errorf("Result is invalid: returned: %v, expected: %v", result, "<b>BIR</b>")
	}
}
