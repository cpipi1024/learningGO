package array

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if want != got {
			t.Errorf("got %d want %d the numbers is '%v'", got, want, numbers)
		}
	})

	t.Run("collections 3", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if want != got {
			t.Errorf("got %d want %d the numbers is '%v'", got, want, numbers)
		}
	})
}
