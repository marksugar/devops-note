package etcdtask

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
	Mark        string `json:"mark"`
	CreatedTime string `json:"create_time"`
	UpdatedTime string `json:"update_time"`
}

const (
	taskKeyPrefix = "/tasks/"
)

type test struct {
	client *clientv3.Client
}

func NewEtcdClient() (test, error) {
	config := clientv3.Config{
		Endpoints:   []string{"172.25.110.31:2379"}, // etcd endpoints
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		return test{}, err
	}
	return test{client: client}, nil
}
func (c test) Close() {
	c.client.Close()
}
func (c test) AddTask(task Task) error {
	id := uuid.New().String()
	now := time.Now().Format(time.RFC3339)
	key := fmt.Sprintf("%s%s", taskKeyPrefix, id)

	record := Task{
		ID:          id,
		Name:        task.Name,
		Status:      task.Status,
		Mark:        task.Mark,
		CreatedTime: now,
		UpdatedTime: now,
	}

	data, err := json.Marshal(record)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Task added: %v,%s\n", string(data), key)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = c.client.Put(ctx, key, string(data))
	cancel()
	if err != nil {
		fmt.Println("put", err)
		return err
	}

	return err
}

func (c test) GetTask(id string) (*Task, error) {
	key := fmt.Sprintf("%s%s", taskKeyPrefix, id)
	fmt.Println(key)

	resp, err := c.client.Get(context.Background(), key)
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, fmt.Errorf("task not found")
	}
	var task Task
	err = json.Unmarshal(resp.Kvs[0].Value, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (c test) UpdateTask(task Task) error {
	key := fmt.Sprintf("%s%s", taskKeyPrefix, task.ID)

	now := time.Now().Format(time.RFC3339)
	task.UpdatedTime = now
	taskData, err := json.Marshal(task)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = c.client.Put(ctx, key, string(taskData))
	cancel()
	return err
}

func (c test) DeleteTask(id int) error {
	key := fmt.Sprintf("%s%d", taskKeyPrefix, id)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := c.client.Delete(ctx, key)
	cancel()
	return err
}

func (c test) ListTasks() ([]*Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := c.client.Get(ctx, taskKeyPrefix, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return nil, err
	}
	var tasks []*Task
	for _, kv := range resp.Kvs {
		var task Task
		err := json.Unmarshal(kv.Value, &task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (c test) WatchTasks() {
	rch := c.client.Watch(context.Background(), taskKeyPrefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			var task Task
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
