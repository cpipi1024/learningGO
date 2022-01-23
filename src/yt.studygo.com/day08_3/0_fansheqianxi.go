package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a float32 = 3.14

	fmt.Printf("类型对象Value对象 : %v\n", reflect.ValueOf(a))
	fmt.Printf("等价于: %v", reflect.ValueOf(a).Float())
}
