package main

import (
	"fmt"
	"log/slog"
	"redis/active"
	"redis/models"
	"redis/service"
)

func main() {
	if err := active.Init(); err != nil {
		slog.Error(fmt.Sprintf("V%", err))
	}
	defer active.Close()

	msgChan := make(chan models.TaskMessage)

	go active.TaskChannel2(msgChan)
	go func() {
		for op := range msgChan {
			for _, ds := range op.Data {
				fmt.Println("[TaskChannel tasks]:", op, "[UUID]:", ds.ID)
				if op.Operation == models.Delete || ds.Status != 1 {
					fmt.Println(" RemoveJob tasks:", ds.ID, ds.Status)
					continue
				}
				if op.Operation == models.Update {
					fmt.Println(" UpdateScheduleTasks tasks:", ds.ID, ds.Status)
				}
				if op.Operation == models.Add {
					fmt.Println(" ScheduleTasks tasks:", ds.ID, ds.Status)
				}
			}
		}
	}()

	a := service.Test{}
	go a.RunData()

	select {}
}
