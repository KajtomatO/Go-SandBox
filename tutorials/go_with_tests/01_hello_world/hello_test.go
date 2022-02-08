package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		//Helpers role is to inform test suite that this IS a helper :-) so fails should be reported when it is called and not inside its body.
		//So, inside a test and not inside a helper function
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Zbigniew", "")
		want := "Hello, Zbigniew!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {

		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {

		got := Hello("Ala", "French")
		want := "Bonjour, Ala!"

		assertCorrectMessage(t, got, want)
	})
}
