package slicetest

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// go中对slice不能使用等号运算符，可以使用一个迭代来检查每一个元素。相对较简单是方法是使用reflect.DeepEqual
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
