package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 使用通道完成完成阻塞
type newlogger interface {
	logging(chan int)
}

func startlogging(logger newlogger, ch1 chan int) {
	logger.logging(ch1)
}

type filelogger struct {
	filename string
}

func (f *filelogger) logging(ch chan int) {
	// 判断文件是否存在
	if file_info, _ := os.Stat(f.filename); file_info == nil {
		file, err := os.Create(f.filename)
		if err != nil {
			panic("创建文件出现问题")
		}
		for i := 0; i < 10; i++ {
			/* if i%5 == 0 {
				time.Sleep(time.Second * 1)
			} */
			file.Write([]byte(time.Now().String() + "\n")) //在文件尾增量写
			num := <-ch
			fmt.Printf("从管道中接收数据: %d\n", num)
		}

	} else {
		file, err := os.OpenFile(f.filename, os.O_APPEND, 0666)
		if err != nil {
			panic("打开文件出现问题")
		}
		for i := 0; i < 10; i++ {
			/* if i%5 == 0 {
				time.Sleep(time.Second * 1)
			} */
			file.Write([]byte(time.Now().String() + "\n")) //在文件尾增量写
			num := <-ch
			fmt.Printf("从管道中接收数据: %d\n", num)
		}
	}
	//close(*ch)
}

type terminallogger struct {
}

func (t terminallogger) logging(ch chan int) {
	for i := 0; i < 10; i++ {
		/* if i%5 == 0 {
			time.Sleep(time.Second * 1)
		} */
		io.WriteString(os.Stdout, time.Now().String()+"\n")
		fmt.Printf("向管道中发送数据: %d\n", i)
		ch <- i // 向管道发送数据
	}

}

/* func main() {
	ch1 := make(chan int)
	fl := filelogger{filename: "1_file.log"}
	tl := terminallogger{}
	go startlogging(tl, ch1)
	go startlogging(&fl, ch1)
	select {
	case <-ch1:
		close(ch1)
		fmt.Printf("完成case!\n")
	}
}
*/
