package active

import (
	"context"
	"encoding/json"
	"etcd/models"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var Ed *clientv3.Client

type etcd struct {
	client *clientv3.Client
}

func Initdb() (etcd, error) {
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // etcd endpoints
		DialTimeout: 5 * time.Second,
	}

	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	Ed = client
	return etcd{client: client}, err
}

func (e etcd) Close() {
	e.client.Close()
}
func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (c etcd) WatchTasks() {
	rch := c.client.Watch(context.Background(), taskKeyPrefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			var task models.Tasks
			if ev.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(ev.Kv.Value, &task)
				if err != nil {
					log.Printf("Failed to unmarshal task: %v", err)
					continue
				}
			}
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Printf("Task added/updated: %v\n", task)
			case clientv3.EventTypeDelete:
				fmt.Printf("Task deleted: %s\n", ev.Kv.Key)
			}
		}
	}
}
func (c etcd) TaskChannel2(msgChan chan<- models.Tasks) {
	rch := c.client.Watch(context.Background(), taskKeyPrefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			var message models.Tasks
			if ev.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(ev.Kv.Value, &message)
				if err != nil {
					log.Printf("Failed to unmarshal task: %v", err)
					continue
				}
			}
			switch ev.Type {
			case clientv3.EventTypePut:
				msgChan <- message
			case clientv3.EventTypeDelete:
				msgChan <- message
			}
		}
	}
}
