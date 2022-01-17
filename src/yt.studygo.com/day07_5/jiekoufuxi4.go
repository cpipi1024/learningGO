package main

import (
	"fmt"
	"os"
)

// 接口练习题 实现一个既可以往终端写日志也可以往文件写日志的简易日志库
type Writer interface {
	write() //往终端和文件写的类型都实现了Writer接口
}

//终端日志
type terminalwrite struct{}

//文件日志
type filewrite struct{}

func (tw terminalwrite) write() {
	fmt.Println("正在向终端写日志")
	//return false
}

func (fw filewrite) write() {
	//return false
	/* pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("报错了！")
	}
	fmt.Println(pwd) */
	msg := "正在向文件中写日志"
	var file *os.File
	file, _ = os.Create("log.log")
	file.WriteString(msg)
	file.Close() // 关闭写入流
	fmt.Println("写入完成")
}
func main() {
	var writer Writer
	var input int
	fmt.Scanln(&input)
	//input = 1 往终端写日志
	//input = 2 往文件写日志
	switch input {
	case 1:
		writer = terminalwrite{}
		writer.write()
	case 2:
		writer = filewrite{}
		writer.write()
	}
}
