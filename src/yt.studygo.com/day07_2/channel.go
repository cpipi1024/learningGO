package main

import (
	"fmt"
)

type stringnumbers interface {
	string()
}

type numbers int

func (n numbers) string() string {
	return fmt.Sprintf("%d", n)
}

// 通道是连接多个协程通信的管道 可以从一个go协程将值发送到通道， 然后在其他协程中接受
func main() {
	//messages := make(chan string)
	// 缓存通道 缓存数量为2
	messages := make(chan string, 2)

	// 调用协程发送信息到chanle中

	go func() {
		/* 		time.Sleep(1 * time.Second)
		   		messages <- "ping" */
		for i := 2; i > 0; i-- {
			messages <- fmt.Sprint(i)
		}
	}()

	for i := 0; i < 2; i++ {
		msg := <-messages
		fmt.Println("msg in chanle: ", msg)
	}

}
