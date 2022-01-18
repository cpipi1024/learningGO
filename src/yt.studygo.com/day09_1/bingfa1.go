package main

import (
	"sync"
)

// 存在多个goroutine同时操作一个资源，这种个情况会发生竞态问腿
var x int // 临界区
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 125000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}

/* func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()

	fmt.Println(x)
}
*/
