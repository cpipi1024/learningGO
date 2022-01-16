package main

import (
	"fmt"
	"time"
)

// 协程学习 go by example

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "times: ", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 正常执行函数f
	var s string
	s = "hello"
	f(s)

	// 采用go 关键字启动协程
	go f("routine")

	// 启动匿名协程
	go func(word string) {
		fmt.Println(word)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done")
}
