package main

import "testing"

func TestHello(t *testing.T) {
	// testing.TB can call helper functions from a Test, or a Benchmark
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sarah", "English")
		want := "Hello, Sarah"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello to people in Spanish", func(t *testing.T) {
		got := Hello("Sarah", "Spanish")
		want := "Hola, Sarah"

		assertCorrectMessage(t, got, want)
	})
}