package dictionary

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dict := map[string]string{
		"test": "this is a test",
	}
	got := Search(dict, "test")
	want := "this is a test"

	if want != got {
		t.Errorf("got: %s want: %s", got, want)
	}
}
