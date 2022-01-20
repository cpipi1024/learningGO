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

// 这里的方法名首字母必须大写
func (s Student) Study() string {
	msg := "好好学习 天天向上"
	fmt.Println(msg)
	return msg
}

// 这里的方法名首字母必须大写
func (s Student) Sleep() string {
	msg := "好好睡觉 精神好好"
	fmt.Println(msg)
	return msg
}

func printMethod(a interface{}) {
	t := reflect.TypeOf(a)  // 返回 Type值类型
	v := reflect.ValueOf(a) // 返回 Value值类型

	// 结构体的方法个数
	fmt.Println("方法个数:", t.NumMethod())

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name) //reflec.Type.Mthod() 返回 Method， Method中的Name字段就存储了函数名信息
		fmt.Printf("method:%s\n", methodType)            // Value类型值代表运行时的数据
		var args = []reflect.Value{}
		v.Method(i).Call(args) // 通过反射调用方法的参数必须是[]reflect.Value类型
	}
}

/* func main() {
s1 := Student{
	Name:  "yangtao",
	Age:   24,
	Score: 99,
}

t := reflect.TypeOf(s1) // 返回类型对象

fmt.Printf("type: %v kin: %v\n", t.Name(), t.Kind()) */

// 通过索引获取指定的结构字段信息
/* 	for i := 0; i < t.NumField(); i++ {
	field := t.Field(i)
	fmt.Printf("name:%s index: %d type: %v json tag: %v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
} */

// 通过字段名获取指定结构体字段信息
/* 	if socrefield, ok := t.FieldByName("Score"); ok {
	fmt.Printf("name: %s index: %d type: %v json tag: %v\n", socrefield.Name, socrefield.Index, socrefield.Type, socrefield.Tag.Get("json"))
} */

// printMethod
/* 	printMethod(s1)

} */
