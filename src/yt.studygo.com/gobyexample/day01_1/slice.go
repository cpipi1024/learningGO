package main

import "fmt"

func main() {
	// 不像数组 slice的类型仅由它包含的元素决定（不像数组中还需要元素的个数）
	s := make([]string, 3)
	fmt.Println("empty :", s)

	//可以像数组一样获取和设置值
	s[0] = "Hello"
	s[1] = "World"
	s[2] = "!"
	fmt.Println("set: ", s)

	// 作为数组的补充，支持比数组更多的操作 例如append 返回一个或多个新的slice
	s1 := append(s, "Hola") // 新的slice
	fmt.Println("s1: ", s1)
	s2 := append(s, "Bonjour")

	fmt.Println("s2: ", s2)

	// slice可以被复制
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	// 在一行代码中声明并初始化slice
	l := []string{"hello", "nmsl!"}
	fmt.Println("l: ", l)

	l1 := make([]int, 3)
	//l1[4] = 4 不能这样复制
	fmt.Println("l1: ", l1)

	// 多维切片
	towD := make([][]int, 3) // 创建的是最外层空间
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		towD[i] = make([]int, innerlen) // 创建内层空间
		for j := 0; j < innerlen; j++ {
			towD[i][j] = j
		}
	}

	fmt.Println("towD: ", towD)
}
