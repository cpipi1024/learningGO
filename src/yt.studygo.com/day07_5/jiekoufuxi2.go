package main

import "fmt"

// 空接口是指没有定义任何方法的接口。 因此所有类型都实现了空接口 ！！！

func show(a interface{}) {
	// 参数可以接受任意类型
	fmt.Printf("Type: %T\n", a)
}

/* func main() {
	s := "string"
	i := 1
	arr := [1]int{1}
	slice := []int{}
	dic := map[string]int{}

	show(s)
	show(i)
	show(arr)
	show(slice)
	show(dic)

	// 空接口还能作为值

	dic1 := make(map[string]interface{})
	dic1["a"] = "hello"
	dic1["b"] = 1
	dic1["c"] = [2]int{1, 2}
	fmt.Println("dic1: ", dic1)
} */
