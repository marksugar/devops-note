package main

import (
	"fmt"
	"time"
)

func main() {

	// Select用于监听多个channel的读写操作
	ch1 := make(chan string)

	go func() {
		select {
		case msg1, ok := <-ch1:
			if !ok {
				fmt.Println("not ok")
				return
			}
			fmt.Println(msg1)
		default:
			fmt.Println("select default")
		}

		fmt.Println("结束")
	}()

	// close(ch1) // 如果close 则会打印not ok
	// ch1 <- "hello select" // 如果写入ch1 <- "hello select"则会打印"hello select"

	// 如果什么都没有，则会等待。如果存在default则进入default

	time.Sleep(1 * time.Second)
	fmt.Println("程序结束")

	many2()
}

// 多通道case同时满足时候是随机执行的
func many2() {

	// Select用于监听多个channel的读写操作
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 2)
	go func() {
		for {
			select {
			case msg1, ok := <-ch1:
				if !ok {
					fmt.Println("not ok")
					return
				}
				fmt.Println(msg1)
			case msg2, ok := <-ch2:
				if !ok {
					fmt.Println("not ok")
					return
				}
				fmt.Println(msg2)
			}
		}
	}()

	// close(ch1) // 如果close 则会打印not ok
	ch1 <- "hello1 select" // 如果写入ch1 <- "hello select"则会打印"hello select"
	ch2 <- "hello2 select"
	// 如果什么都没有，则会等待。如果存在default则进入default

	time.Sleep(1 * time.Second)
	fmt.Println("程序结束")
}
