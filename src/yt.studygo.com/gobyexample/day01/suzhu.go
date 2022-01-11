package main

import "fmt"

func main() {
	// 数组里面的元素类型是一致的
	var a [5]int
	fmt.Println("empty: ", a)
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a)

	var s [5]string

	fmt.Println("empty string array: ", s)
	s[1] = "hello"
	fmt.Println("not empty string array: ", s)

	b := [4]int{1, 2, 3, 4}
	fmt.Println("b: ", b)

	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = j
		}
	}
	fmt.Println("2D array: ", twoD)
}
