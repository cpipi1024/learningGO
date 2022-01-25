package main

import (
	"fmt"
	"math/rand"
)

// 使用goroutine和channel实现一个计算int64随机数各位数和的程序。
func worker_makenum(receive chan<- int64, num_rand rand.Rand) {

	for i := 0; i < 100; i++ {
		num := num_rand.Int63()
		receive <- num
	}
	close(receive)
}

func calculate(receive chan<- int64, send <-chan int64) {

	for num := range send {
		var res int64 = 0
		for {
			temp := num % 10
			res += temp
			num = num / 10
			if num < 1 {
				break
			}
		}
		receive <- res
	}
	close(receive)
}

func main() {
	var j int = 0
	numchan := make(chan int64, 200) // 输入数据通道
	reschan := make(chan int64, 100) // 结果通道

	source := rand.NewSource(10)
	num_rand := rand.New(source)
	/* 	for i := 0; i < 100; i++ {
	   		worker_makenum(numchan, *num_rand)
	   	}
	   	close(numchan) */
	//go worker_makenum(numchan, *num_rand)
	for i := 0; i < 24; i++ {
		go calculate(reschan, numchan)
	}
	/* for num := range numchan {
		fmt.Println(num)
	} */
	go worker_makenum(numchan, *num_rand)
	//close(reschan)
	for res := range reschan {
		fmt.Printf("res:%d\n", res)
		j += 1
	}
	fmt.Println("j:", j)
}
