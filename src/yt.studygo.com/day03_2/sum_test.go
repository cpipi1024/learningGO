package array

import "testing"

func TestSum(t *testing.T) {
	var numbers [5]int
	numbers = [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if want != got {
		t.Errorf("got %d want %d number is %v", got, want, numbers)
	}
}
