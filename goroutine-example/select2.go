package main

import (
	"fmt"
	"time"
)

// 消息收发
// 超时机制
// 定时器

func main() {
	// 消息收发
	serverch := make(chan string, 10)
	clientch := make(chan string, 10)

	// 服务器端
	go func() {
		for {
			select {
			case msg, ok := <-serverch:
				if !ok {
					fmt.Println("serverch closed")
					break
				}
				fmt.Printf("serverch read: %s \n", msg)
				clientch <- fmt.Sprintf("-> reply to: -> %s", msg)
			}
		}
	}()
	// 客户端
	go func() {
		for {
			select {
			case msg, ok := <-clientch:
				if !ok {
					fmt.Println("clientch closed")
					break
				}
				fmt.Printf("clientch read:%s \n ", msg)
			}
		}
	}()

	serverch <- "hello world, i am client!"
	time.Sleep(2 * time.Second)
}
