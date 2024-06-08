package main

import (
	"fmt"
	"os"
	"time"

	"github.com/percona/go-mysql/log"
	"github.com/percona/go-mysql/log/slow"
)

func main() {
	// 打开慢查询日志文件
	filePath := "slow.log"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 获取文件的初始大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}
	initialSize := fileInfo.Size()

	// 创建慢查询日志解析器
	options := log.Options{Debug: false}
	parser := slow.NewSlowLogParser(file, options)

	// 通道，用于处理事件
	eventChan := make(chan log.Event)
	doneChan := make(chan struct{})

	// 启动监控文件的 Goroutine
	go monitorFile(filePath, parser, initialSize, eventChan, doneChan)

	// 处理解析后的事件
	for event := range eventChan {
		processEvent(event)
	}

	// 等待监控 Goroutine 结束
	<-doneChan
}

func monitorFile(filePath string, parser *slow.SlowLogParser, initialSize int64, eventChan chan<- log.Event, doneChan chan<- struct{}) {
	defer close(doneChan)
	defer close(eventChan)

	for {
		// 打开文件
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error reopening file: %v\n", err)
			return
		}

		// 获取当前文件大小
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			file.Close()
			continue
		}

		// 如果文件大小变化，读取新内容
		if fileInfo.Size() > initialSize {
			if _, err := file.Seek(initialSize, 0); err != nil {
				fmt.Printf("Error seeking file: %v\n", err)
				file.Close()
				continue
			}

			// 逐行读取新内容
			parser := slow.NewSlowLogParser(file, log.Options{Debug: false})
			go parser.Start()

			for event := range parser.EventChan() {
				eventChan <- *event
			}

			// 更新已读取的文件大小
			initialSize = fileInfo.Size()
		}

		file.Close()

		// 每隔一段时间检查一次文件变化
		time.Sleep(1 * time.Second)
	}
}

func processEvent(event log.Event) {
	// 检查Query_time时间是否超过 5 秒
	if event.TimeMetrics["Query_time"] > 5 {
		fmt.Println("db:", event.Db, "host:", event.Host, "user:", event.User, "time:", event.Ts, "local/query-time:", event.TimeMetrics)
		fmt.Println("\n", event.Query)
	}
}
