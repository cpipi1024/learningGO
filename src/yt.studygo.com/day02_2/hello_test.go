package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		t.Helper() // 告诉测试套件当前这个方法是辅助函数，当报错时报错的行数就会在函数调用中而不是辅助函数中
		if got != want {
			t.Errorf("got: '%q' want: '%q'", got, want)
		}
	}
	// 测试用例 1
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("chris", "")
		want := "Hello, chris"
		assertMessage(t, got, want)
	})
	// 测试用例 2
	t.Run("saying hello to empty", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, world"
		assertMessage(t, got, want)
	})
	// 测试用例 3
	t.Run("saying in spanish", func(t *testing.T) {
		got := hello("elodie", "spanish")
		want := "Hola, elodie"
		assertMessage(t, got, want)
	})
	// 测试用例 4
	t.Run("saying in franch", func(t *testing.T) {
		got := hello("elodie", "franch")
		want := "Bonjour, elodie"
		assertMessage(t, got, want)
	})
}
