package main

//func main() {
//	//选择器可以做到同时等待多个通道
//	c1 := make(chan string)
//	c2 := make(chan string)
//
//	go func() {
//		time.Sleep(time.Second)
//		c1 <- "channel 1"
//	}()
//
//	go func() {
//		time.Sleep(time.Second * 2)
//		c2 <- "channel 2"
//	}()
//	// 输入会阻塞整个程序
//	/*
//		var input string
//		   	fmt.Scanln(&input)
//	*/
//	for i := 0; i < 2; i++ {
//		select {
//		case msg1 := <-c1:
//			fmt.Println("received: ", msg1)
//		case msg2 := <-c2:
//			fmt.Println("received: ", msg2)
//		}
//	}
//}
