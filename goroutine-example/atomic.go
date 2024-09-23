package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 除了加锁来保证并发安全之外，也可以利用原子操作保证并发安全
// 相比加锁操作，原子操作可以减少资源的消耗
var i int64

var swg sync.WaitGroup

// 普通
func incre() {
	i = i + 1
	swg.Done()
}

// Mutex加锁操作
var lock sync.Mutex

func mutexIncre() {
	lock.Lock()
	defer lock.Unlock()

	i = i + 1
	swg.Done()
}

// 原子操作
func atomicIncre() {
	defer swg.Done()
	atomic.AddInt64(&i, 1)
}

func main() {
	for i := 0; i < 1000000; i++ {
		swg.Add(1)
		// go mutexIncre()  // 加锁
		// go incre() 		// 普通
		go atomicIncre() // 原子操作
	}
	swg.Wait()
	fmt.Println(i)
}
