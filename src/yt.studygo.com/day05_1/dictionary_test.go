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
	t.Run("添加不存在的键值", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("unknown", "this is also a test")
		got, err := dict.Search("unknown")
		want := "this is also a test"
		if err != nil {
			t.Fatal("should found a added value: ", err)
		}

		if got != want {
			t.Errorf("got : %s want: %s", got, want)
		}
	})
	t.Run("添加已存在键值", func(t *testing.T) {
		//dict := Dictionary{}
		word := "test"
		definition := "this is a new test!"

		dict := Dictionary{word: definition}

		err := dict.Add(word, "new test!")
		assertError(t, err, ErrorExists)
		assertDefinition(t, dict, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("更新存在的键值", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := Dictionary{word: definition}

		new_definition := "new test"

		dict.Update(word, new_definition)

		assertDefinition(t, dict, word, new_definition)
	})

	t.Run("更新不存在的键值", func(t *testing.T) {
		word := "unknown"
		definition := "unknown value "
		dict := Dictionary{}

		err := dict.Update(word, definition)
		assertError(t, err, ErrorKeyNotExists)
		//assertDefinition(t, dict, word, definition)
	})
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()

	if want != got {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) { //为啥这里的got形参不用定义类型？ 参数两个一样类型可以指声明一次
	t.Helper()

	if want != got {
		t.Errorf("got error: %s want error: %s", got, want)
	}
}

// 新的断言 为已存在键的map重新赋值

func assertDefinition(t *testing.T, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should found a added value: ", err)
	}
	if definition != got {
		t.Errorf("got: %s want: %s", got, definition)
	}

}
