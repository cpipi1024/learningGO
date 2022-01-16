package main

// 非阻塞通道通信
/* func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received messages: ", msg)
	default:
		fmt.Println("no messages received")
	}

	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("send messages: ", msg)
	default:
		fmt.Println("no message sent")
	}

	// 多路非阻塞的选择器
	select {
	case msg := <-messages:
		fmt.Println("received messages: ", msg)
	case sig := <-signals:
		fmt.Println("received signals: ", sig)
	default:
		fmt.Println("no activity")
	}
} */
