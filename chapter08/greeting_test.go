package main

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buf := bytes.Buffer{}
	name := "uniyuni1101"
	want := "Hello, uniyuni1101"

	Greet(&buf, name)
	got := buf.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
