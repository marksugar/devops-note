package main

import (
	"fmt"
	"tgsed/svc"
)

func main() {
	// Initialize the bot service
	b := svc.TgbotService.TgCallbackInit()

	// Use the service to send a message
	err := b.SendMessageAlone(-4215, "这是一个测试")
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
