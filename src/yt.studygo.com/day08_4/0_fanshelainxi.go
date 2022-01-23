package main

import (
	"fmt"
	"reflect"
)

type phone int64
type person struct {
	name     string
	age      int
	addr     string
	p_number phone
}

// 不含参数的方法
func (p person) Speak() {
	fmt.Printf("我是: %s, 今年: %d岁, 住在: %s, 电话号码是: %d\n", p.name, p.age, p.addr, p.p_number)
}

// 含有参数的方法
func (p person) Work(job string) {
	fmt.Printf("我做的工作是: %s\n", job)
}

func reflectFunc(a interface{}) {
	rType := reflect.TypeOf(a)
	rValue := reflect.ValueOf(a)

	fmt.Printf("person结构体的concreate类型: %v, static类型: %v\n", rType, rType.Kind())
	// 打印结构体字段信息
	for i := 0; i < rType.NumField(); i++ {
		field_name := rType.Field(i).Name
		field_type := rType.Field(i).Type.Name()
		fmt.Printf("person结构体中的第%d字段: %s\t", i+1, field_name)
		fmt.Printf("person结构体中的第%d字段类型: %v\n", i+1, field_type)
	}
	// 打印结构体方法信息
	for j := 0; j < rType.NumMethod(); j++ {
		method_name := rType.Method(j).Name
		method_type := rType.Method(j).Type
		method_value := rType.Method(j).Func
		fmt.Printf("person结构体中的第%d个方法的方法名: %s\n", j+1, method_name)
		fmt.Printf("person结构体中的第%d个方法的方法类型: %v\n", j+1, method_type)
		fmt.Printf("person结构体中的第%d个方法的方法的值: %v\n", j+1, method_value)
	}
	// 尝试调用p中的方法
	for k := 0; k < rType.NumMethod(); k++ {
		// 根据参数是否有的情况调用
		invoke_method := rValue.Method(k) // 返回v持有的值类型的第k个方法的已绑定状态函数形式的Value封装
		fmt.Printf("方法的Value封装: %v\n", invoke_method)
		fmt.Printf("方法中的参数: %d\n", invoke_method.Type().NumIn())
		switch invoke_method.Type().NumIn() {
		case 0:
			param := []reflect.Value{}
			invoke_method.Call(param)
		case 1:
			param := []reflect.Value{reflect.ValueOf("搬砖")}
			invoke_method.Call(param)
		}
	}
}

func main() {
	p := person{
		name:     "杨涛",
		age:      24,
		addr:     "深圳",
		p_number: 1485764284,
	}

	reflectFunc(p)
	//p.speak()
}
