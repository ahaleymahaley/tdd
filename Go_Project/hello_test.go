package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lena")
	want := "Hello, Lena!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
