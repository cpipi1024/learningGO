package main

import (
	"bytes"
	"testing"
)

func TestCuntDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spysleeper := &SpySleeper{}

	CountDown(buffer, spysleeper)

	got := buffer.String()
	want := `3
2
1
go!`

	if want != got {
		t.Errorf("got: %s want: %s", got, want)
	}

	if spysleeper.Calls != 4 {
		t.Errorf("not enough calls to sleep. want 4 got %d", spysleeper.Calls)
	}
}
