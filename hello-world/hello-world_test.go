package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := HelloWorld("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say more personalised 'Hello, <name>' when a name is provided", func(t *testing.T) {
		got := HelloWorld("Manni", "")
		want := "Hello, Manni"
		assertCorrectMessage(t, got, want)

	})
	t.Run("Say Hello (Hola) in Spanish", func(t *testing.T) {
		got := HelloWorld("Fernando", "Spanish")
		want := "Hola, Fernando"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello (Bonjour) in French", func(t *testing.T) {
		got := HelloWorld("Marie", "French")
		want := "Bonjour, Marie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello (Sat Sri Akal) in Punjabi", func(t *testing.T) {
		got := HelloWorld("Diljit", "Punjabi")
		want := "Sat Sri Akal, Diljit"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
