package main

import (
	"fmt"
	"time"
)

func timer() {
	// 延迟处理
	timer := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	<-timer.C // 等待延迟结束

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
func main() {
	// 定时器,计划任务

	// 秒
	t1 := time.Tick(time.Second)
	t3 := time.Tick(time.Second * 3)
	t5 := time.Tick(time.Second * 5)

	for {
		select {
		case <-t1:
			fmt.Println("t1 sec scheduler")
		case <-t3:
			fmt.Println("t3 sec scheduler")
		case <-t5:
			fmt.Println("t5 sec scheduler")
		}
	}
}
