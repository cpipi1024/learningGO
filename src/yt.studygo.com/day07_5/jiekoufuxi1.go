package main

import "fmt"

// 接口中的方法不一定需要全部实现 可以在类型中嵌入其他类型或者结构体

type WashMachine interface {
	wash()
	dry()
}

type Dryer struct{}

func (d Dryer) dry() {
	fmt.Println("甩干中")
}

type Haier struct {
	Dryer //嵌套的结构体
}

func (h Haier) wash() {
	fmt.Println("洗刷刷洗刷刷")
}

/* func main() {
	h := Haier{}
	h.wash()
	h.dry()
} */
