package main

import "testing"

func TestName(t *testing.T) {
	name := getName()
	RESULT := "World"
	if name != RESULT {
		t.Error("Expected", RESULT, "Got", name)
	}
}
