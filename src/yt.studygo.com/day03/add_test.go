package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := add(1, 2)
	expect := 3
	if sum != expect {
		t.Errorf("expected '%d' but got '%d'", expect, sum)
	}
}
