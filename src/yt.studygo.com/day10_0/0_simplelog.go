package main

import (
	"io"
	"os"
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
	for i := 0; i < 10; i++ {
		io.WriteString(os.Stdout, time.Now().String()+"\n")
	}
}

type FileLogger struct {
	filename string
}

func (f *FileLogger) loging() {
	fileinfo, _ := os.Stat(f.filename)
	if fileinfo == nil {
		file, err := os.Create(f.filename)
		if err != nil {
			panic("创建文件出错!")
		}
		for i := 0; i < 10; i++ {
			file.WriteString(time.Now().String() + "\n")
		}
	} else {
		file, err := os.Open(f.filename)
		if err != nil {
			panic("打开文件出错!")
		}
		for i := 0; i < 10; i++ {
			file.WriteString(time.Now().String() + "\n")
		}
	}

}

/* func main() {
	fl := FileLogger{
		filename: "file.log",
	}
	tl := TerminalLogger{}
	go fl.loging()
	go tl.loging()
	time.Sleep(time.Second * 5)
	fmt.Println("main go 结束")
}
*/
