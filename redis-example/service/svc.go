package service

import (
	"log"
	"redis/active"
	"redis/models"
	"time"
)

type Test struct {
	at active.SystemTest
}

func (s Test) RunData() {

	// 创建新数据
	newTask := models.Test{
		ID:     1,
		Name:   "New Task",
		Status: 1,
	}

	if err := s.at.PublishTaskChange(active.Rdb, models.OperationType(models.Add), newTask); err != nil {
		log.Fatalf("Failed to publish data change: %v", err)
	}
	var id = 1
	// 模拟数据更新
	go func() {
		for {
			time.Sleep(2 * time.Second)
			var task models.Test
			task.ID = id
			task.Name = "Updated Task"
			if err := s.at.PublishTaskChange(active.Rdb, models.OperationType(models.Add), task); err != nil {
				log.Fatalf("Failed to publish data change: %v", err)
			}
			id++
		}
	}()
}
