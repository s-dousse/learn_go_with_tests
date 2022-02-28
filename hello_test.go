package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sarah")
	want := "Hello, Sarah"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}