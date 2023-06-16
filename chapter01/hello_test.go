package chapter01

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("うにゆに", "Japanese")
		want := "こんにちは、うにゆに"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("うにゆに", "French")
		want := "Bonjour, うにゆに"
		assertCorrectMessage(t, got, want)
	})

	t.Run("unregistered language", func(t *testing.T) {
		got := Hello("うにゆに", "Chinese")
		want := "Hello, うにゆに"
		assertCorrectMessage(t, got, want)
	})
}
