package main

import (
	"fmt"
	"sync"
	"time"
)

// 加锁
// 加锁后其他进程不能操作他,避免脏读

// 定义一个secKill函数，调用secKill的时候会从remain变量减去1，如果remain小于等于0则返回，不在相减
// main中定义一个for循环，循环小于10的数，并且使用wg，每次循环加1，等待调用完成
// 同时在secKill函数中，每个调用会使用mu定义的互斥锁锁定，在结束后释放mu.Unlock()，并且结束协程wg.Done()
// 当remain的值被使用完后就结束协程，不会出现负数

var remain = 5
var mu sync.Mutex // 定义互斥锁
var wg sync.WaitGroup

func secKill() {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	if remain <= 0 {
		return
	}
	time.Sleep(100 * time.Millisecond)
	remain = remain - 1
	fmt.Println("seckill success")
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go secKill()
	}
	wg.Wait()
	fmt.Println("======done======")
}
