package main

import (
	"fmt"
	"reflect"
)

//结构体反射
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"socre"`
}

func (s Student) study() string {
	msg := "好好学习 天天向上"
	fmt.Println()
	return msg
}

func (s Student) sleep() string {
	msg := "好好睡觉 精神好好"
	fmt.Println()
	return msg
}

func printMethod(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	// 结构体的方法个数
	fmt.Println(t.NumMethod())
	
	for i := 0; i < v.NumMethod(); i++ {
		methodType := 
	}
}

func main() {
	s1 := Student{
		Name:  "yangtao",
		Age:   24,
		Score: 99,
	}

	t := reflect.TypeOf(s1) // 返回类型对象

	fmt.Printf("type: %v kin: %v\n", t.Name(), t.Kind())

	// 通过索引获取指定的结构字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index: %d type: %v json tag: %v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if socrefield, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name: %s index: %d type: %v json tag: %v\n", socrefield.Name, socrefield.Index, socrefield.Type, socrefield.Tag.Get("json"))
	}

}
