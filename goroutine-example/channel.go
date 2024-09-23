package main

import (
	"fmt"
	"time"
)

// 使用channel
// 除了使用互斥锁解析并发读写之外还可以使用channel，channel用于i数据传递的管道

func receiver(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}
func main() {
	// 10是通道的数量，如果等于for循环的数值则不会阻塞，如果超过则会阻塞，等待通道的数据被取出，而后继续写入
	ch := make(chan string, 10)
	go receiver(ch)

	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("hello %d", i)
	}
	fmt.Println("hello")

	time.Sleep(1 * time.Second)

}
