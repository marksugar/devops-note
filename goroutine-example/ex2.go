package main

import (
	"fmt"
	"sync"
)

// 1,waitgroup用于等待携程运行结束
// 2,多协程同时操作一个变量，会产生脏读，脏写问题等，保证并发的安全使用方式：
//
//	2.1 加锁
//	2.2 channel通信
//
// 3，协程执行的时候会调用add加1，结束后调用defer wg.Done结束，并且-1。而后wg.Wait()等待协程全部结束
func wg() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("======done======")
}

func main() {
	wg()
}
