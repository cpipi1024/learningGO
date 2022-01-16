package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received jobs: ", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		time.Sleep(time.Second)
		jobs <- j
		//time.Sleep(time.Second * 1)
		fmt.Println("send job", j)
	}
	close(jobs)
	fmt.Println("send all jobs")

	//<-done
}
