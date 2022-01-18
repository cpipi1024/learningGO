package main

// 单纯地让函数并发执行是没有意义的。 函数与函数之间需要交换数据才能体现并发执行函数的意义
// go语言的并发模型是提倡通信共享内存 而不是 通过共享内存而实现通信
// 可以通过共享内存实现数据之间的交换， 但是容易出现竞态问题。 为了保证数据交换的正确性，必须使用互斥量对内存加锁，这种做法势必造成性能问题。

// 通道channel 遵循FIFO规则，每个通道都是一具体类型的管道。

/* func main() {
	var a chan int
	fmt.Println(a) // nil chan是引用类型

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("向ch1中发送数据: ", i)
			ch1 <- i
		}
		close(ch1) // 关闭管道
	}()

	go func() {
		for {
			i, ok := <-ch1
			fmt.Println("接收到ch1中发送的数据: ", i)
			if !ok {
				fmt.Println("管道ch1为空")
				break
			}
			fmt.Println("开始向ch2中发送数据: ", i*i)
			ch2 <- i * i
		}
		close(ch2) // 关闭管道
	}()
	for i := range ch2 {
		fmt.Println("ch2管道中存储的值: ", i)
	}

} */
