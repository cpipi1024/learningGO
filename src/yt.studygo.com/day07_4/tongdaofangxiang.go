package main

//指定通道是否用来

func ping(resive chan<- string, msg string) {
	resive <- msg
}

func pong(send <-chan string, resive chan<- string) {
	msg := <-send
	resive <- msg
}

/* func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "sending msg")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
*/
