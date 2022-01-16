package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type Dog struct {
	name string
}

func (d Dog) say() {
	fmt.Printf("%s:汪汪汪\n", d.name)
}

func (d Dog) move() {
	fmt.Printf("%s 在动", d.name)
}

/* func main() {
	// 一个类型可以实现多个接口 并且多个接口之间并不知道相互实现了什么方法
	var d = Dog{name: "旺财"}

	var s Sayer
	var m Mover
	s = d
	m = d
	s.say()
	m.move()
} */
