package dictionary

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dict := Dictionary{
		"test": "this is a test",
	}
	t.Run("有对应键值: ", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is a test"

		assertString(t, got, want)
	})

	// 找不到键值的情况
	t.Run("没有对应键值: ", func(t *testing.T) {
		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("expected a error")
		}
		//assertString(t, err.Error(), ErrorNotFound)
		assertError(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	dict.Add("test", "this is also a test")
	got, err := dict.Search("test")
	want := "this is also a test"
	if err != nil {
		t.Fatal("should find a added value: ", err)
	}

	if got != want {
		t.Errorf("got : %s want: %s", got, want)
	}
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()

	if want != got {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) { //为啥这里的got形参不用定义类型？
	t.Helper()

	if want != got {
		t.Errorf("got error: %s want error: %s", got, want)
	}
}
