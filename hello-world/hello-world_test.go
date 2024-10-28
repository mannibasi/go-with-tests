package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloPersonalised(t *testing.T) {
	got := HelloWorld("Manni")
	want := "Hello, Manni"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
