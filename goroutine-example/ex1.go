package main

import (
	"fmt"
	"time"
)

// 1.不等待携程
// 2,无序
func main() {
	go fmt.Println("================")
	for i := 0; i < 5; i++ {
		go fmt.Println("ok", i)
	}

	time.Sleep(1 * time.Second)
}
