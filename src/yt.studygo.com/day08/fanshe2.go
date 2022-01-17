package main

import (
	"fmt"
	"reflect"
)

func reflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	kind := v.Kind() //返回的是reflext.Value类型
	switch kind {
	case reflect.Int64:
		fmt.Printf("type is int64 and value is %d\n", int64(v.Int())) //转换reflec.Value类型为对应的类型值
	case reflect.Float32:
		fmt.Printf("type is float32 and value is %f\n", float32(v.Float()))
	}
}

/* func main() {
	var a int64 = 10
	var b float32 = 3.1415

	reflectValue(a)
	reflectValue(b)
} */
