package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 8)
	want := "aaaaaaaa"
	if got != want {
		t.Errorf("got '%q' but want '%q'", got, want)
	}
}

func BenchMarksRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}
