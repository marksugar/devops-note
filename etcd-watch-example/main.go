package main

import (
	"etcd/active"
	"etcd/models"
	"etcd/router"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	e, err := active.Initdb()
	if err != nil {
		panic(err)
	}
	defer e.Close()
	// go e.WatchTasks()
	// go e.WatchTasks()
	msgChan := make(chan models.Tasks)
	go e.TaskChannel2(msgChan)
	go func() {
		for op := range msgChan {
			fmt.Println("[TaskChannel tasks]:", op)
			if op.Operation == models.Delete || op.Status != 1 {
				fmt.Println(" RemoveJob tasks:", op)
				continue
			}
			if op.Operation == models.Update {
				fmt.Println("Update ScheduleTasks tasks:", op)
			}
			if op.Operation == models.Add {
				fmt.Println("Add ScheduleTasks tasks:", op)
			}
			if op.Operation == models.Nothing {
				slog.Info("id update. %v", op.ID)
			}
		}
	}()
	r := router.Setup("dev")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":8080"),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error(fmt.Sprintf("err:%v", err))
	}

	slog.Info(fmt.Sprintf("listen:8080"))

	//

}
