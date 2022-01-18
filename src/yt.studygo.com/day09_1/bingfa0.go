package main

/* func main() {
	// 利用同步通道/无缓存通道 完成同步
	ch1 := make(chan int)
	go func() {
		fmt.Println("接收")
		<-ch1
	}()
	//ch1 <- 2 // 会造成死锁
	ch1 <- 1 // 到这会阻塞
	fmt.Println("main goroutine执行完毕！")
}
*/
