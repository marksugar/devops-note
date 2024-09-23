package main

import (
	"fmt"
	"time"
)

func main() {
	// 超时处理
	ds := make(chan bool)
	go func() {
		// 睡眠5秒后传入true
		time.Sleep(5 * time.Second)
		ds <- true
	}()
	select {
	// 如果ds等于true，打印done
	// 如果事件超过3秒则打印timeout
	case <-ds:
		fmt.Println("done")
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	}
}
