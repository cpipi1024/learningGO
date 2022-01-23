package main

import (
	"fmt"
	"time"
)

// 简单日志库 改良
// 同时写和同时在终端输出
type logger interface {
	loging()
}

type TerminalLogger struct {
}

func (t TerminalLogger) loging() {

}

type FileLogger struct {
}

func (f FileLogger) loging() {

}

func main() {
	fmt.Println(time.Now().String())
}
