package main

import (
	"fmt"
	"reflect"
)

type Myint int64

// 在反射中类型还划分为两种：类型type 和 种类kind
// kind种类 就是指底层数据类型，当需要区分指针、结构体等大品种的类型时，就会用到种类。

func TypeAndKind(a interface{}) {
	v := reflect.TypeOf(a)
	fmt.Printf("type:%v   kind:%v\n", v.Name(), v.Kind())
}

/* func main() {
	var a *float32
	var b Myint
	var c rune

	TypeAndKind(a) // go语言的反射中数组、切片、Map、指针等类型的变量它的Name()都是空
	TypeAndKind(b)
	TypeAndKind(c)

	type Person struct {
		name string
		age  int
	}
	type Book struct {
		title string
	}

	p := Person{
		name: "yangtao",
		age:  12,
	}

	book := Book{title: "golang"}

	TypeAndKind(p) // 类型是定义的person类型 底层数据类型是struct
	TypeAndKind(book)
} */
