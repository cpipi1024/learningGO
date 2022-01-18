package main

import (
	"fmt"
	"time"
)

// 启动goroutine时分配的栈内存是可变化的，在协程的生命周期中开始只有很小的栈,goroutine的栈是不固定的，可以按需增大和缩小

// GPM是go语言运行时层面的实现，是go语言自己实现的一套调度系统。 区别于操作系统调用OS线程

// G很好理解，就是个goroutine，里面除了存放本goroutine信息外，还有与所在P的绑定信息
// P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境，P会对自己管理的goroutine队列做一些调度。
// M是go运行时对操作系统内核线程的虚拟，M与内核线程一般是一一映射的关系，一个goroutine最终是要放到M上执行的

func a() {
	for i := 0; i < 20; i++ {
		fmt.Println("A: ", i)
		if i%10 == 0 && i != 0 {
			time.Sleep(time.Second * 1)
		}
	}
}

func b() {
	for j := 0; j < 10; j++ {
		fmt.Println("B: ", j)
	}
}

/* func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second * 3)
}
*/
