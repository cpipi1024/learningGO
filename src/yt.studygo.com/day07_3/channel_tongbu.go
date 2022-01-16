package main

import (
	"fmt"
	"time"
)

// 通道同步 使用通道来同步go协程之间的执行状态

func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done!")
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)

	res := <-done
	if res == true {
		fmt.Println("执行完成！")
	}
}
