// 依赖注入学习

package main

import (
	"bytes"
	"testing"
)

func TestInjection(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "chris")

	got := buffer.String()
	want := "hello, chris!"

	if got != want {
		t.Errorf("got: %s want: %s", got, want)
	}
}
