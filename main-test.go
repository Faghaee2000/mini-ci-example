package main

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello"
	if Greet() != want {
		t.Fatalf("expected %q, got %q", want, Greet())
	}
}
