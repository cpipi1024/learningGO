package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("chris")
		want := "Hello, chris"
		if got != want {
			t.Errorf("got: '%q' want: '%q'", got, want)
		}
	})
	t.Run("saying hello to empty", func(t *testing.T) {
		got := hello("")
		want := "Hello, world"
		if got != want {
			t.Errorf("got: '%q' want: '%q'", got, want)
		}
	})
}
