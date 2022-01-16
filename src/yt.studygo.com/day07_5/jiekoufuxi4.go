package main

import "fmt"

// 接口练习题 实现一个既可以往终端写日志也可以往文件写日志的简易日志库
type Writer interface {
	write() //往终端和文件写的类型都实现了Writer接口
}

//终端日志
type terminalwrite struct{}

//文件日志
type filewrite struct{}

func (tw terminalwrite) write() {

	//return false
}

func (fw filewrite) write() {
	//return false
}
func main() {
	var writer Writer
	var input int
	fmt.Scanln(&input)
	//input = 1 往终端写日志
	//input = 2 往文件写日志
	switch input {
	case 1:
		writer.write()
	case 2:
		writer.write()
	}
}
