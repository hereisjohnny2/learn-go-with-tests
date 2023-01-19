package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "João")

	got := buffer.String()
	want := "Hello, João"

	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}
