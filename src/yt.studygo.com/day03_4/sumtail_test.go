package slicetest2

import (
	"reflect"
	"testing"
)

func TestSumTail(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
	t.Run("not empty", func(t *testing.T) {
		got := SumTail([]int{0, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("empty ", func(t *testing.T) {
		got := SumTail([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
