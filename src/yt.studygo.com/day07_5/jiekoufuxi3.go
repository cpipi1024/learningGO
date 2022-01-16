package main

import (
	"fmt"
)

// 接口的值
// 涉及内容类型断言
/* func main() {
	// 一个接口的值是由具体类型 和 具体类型的值组成
	// 这两部分可以称为 接口的动态类型和动态值

	justifyType("a")
	justifyType(1)
	justifyType(true)
	justifyType([]int{1, 2, 3})
} */

func justifyType(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("x is string and values is %v\n", v)
	case int:
		fmt.Printf("x is int and value is %v\n", v)
	case bool:
		fmt.Printf("x is bool and value is %v\n", v)
	case []int:
		fmt.Printf("x is slice and value is %v\n", v)
	default:
		fmt.Println("unspport type!")
	}
}
