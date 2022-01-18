package main

import "fmt"

// 单向通道

// 管道方向 out <-
func counter(out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i
	}
	close(out) // 关闭通道
}

//  管道方向 out <- in
// out为接收端  in为发送端
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

// 管道方向 <- in
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

/* func main() {
	ch1 := make(chan int) // out
	ch2 := make(chan int) // in

	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
*/
