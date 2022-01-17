package main

import (
	"fmt"
	"reflect"
)

// 反射就是在运行时动态的获取一个变量的类型信息和值信息

func reflectType(a interface{}) {
	v := reflect.TypeOf(a) // 可以获得任何值的类型对象 通过类型对象可以访问任意值的类型信息
	fmt.Printf("type:%v \n", v)
}

/* func main() {
	var a float32 = 3.14
	reflectType(a)

	var b int64 = 100
	reflectType(b)
} */
