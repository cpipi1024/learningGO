package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("woker: %d, start job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("woker: %d, end job %d\n", id, j)
		results <- j * 2
	}

}

/* func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	start := time.Now()
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		<-results
	}
	cost := time.Now().Sub(start)
	fmt.Printf("cost time: %v", cost)
} */
