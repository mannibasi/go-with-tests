package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := HelloWorld("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Say more personalised 'Hello, <name>' when a name is provided", func(t *testing.T) {
		got := HelloWorld("Manni")
		want := "Hello, Manni"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
