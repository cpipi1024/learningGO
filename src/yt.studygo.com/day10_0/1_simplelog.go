package main

import (
	"io"
	"os"
	"time"
)

// 使用通道完成完成阻塞
type newlogger interface {
	logging()
}

func startlogging(logger newlogger) {
	logger.logging()
}

type filelogger struct {
	filename string
}

func (f *filelogger) logging() {
	// 判断文件是否存在
	if file_info, _ := os.Stat(f.filename); file_info == nil {
		file, err := os.Create(f.filename)
		if err != nil {
			panic("创建文件出现问题")
		}
		for i := 0; i < 10; i++ {
			file.WriteString(time.Now().String() + "\n") //在文件尾增量写
			time.Sleep(time.Second * 1)
		}

	}
	file, err := os.Open(f.filename)
	if err != nil {
		panic("打开文件出现问题")
	}
	for i := 0; i < 10; i++ {
		file.WriteString(time.Now().String() + "\n")
		time.Sleep(time.Second * 1)
	}

}

type terminallogger struct {
}

func (t terminallogger) logging() {
	for i := 0; i < 10; i++ {
		io.WriteString(os.Stdout, time.Now().String()+"\n")
		time.Sleep(time.Second * 1)
	}

}

/* func main() {
	fl := filelogger{filename: "1_file.log"}
	tl := terminallogger{}
	go startlogging(&fl)
	go startlogging(tl)
	select {}

} */
