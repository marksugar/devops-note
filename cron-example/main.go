package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

var s *gocron.Scheduler

type Crons interface {
	Init()
}
type Task struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Interval    int    `json:"interval"` // 以秒为单位
	Func        string `json:"func"`
	Status      int    `json:"status"`
	CreatedTime string `json:"create_time"`
	UpdatedTime string `json:"update_time"`
}
type Cron struct {
	scheduler    gocron.Scheduler
	TaskFuncMaps map[int]GetTask
}

var TaskFuncMaps = map[int]GetTask{
	1: {Func: MyTask1, FuncName: "MyTask1"},
	2: {Func: MyTask2, FuncName: "MyTask2"},
	3: {Func: MyTask3, FuncName: "MyTask3"},
}

func NewCron() (Cron, error) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return Cron{}, err
	}
	s, err := gocron.NewScheduler(gocron.WithLocation(location))
	return Cron{
		scheduler:    s,
		TaskFuncMaps: TaskFuncMaps,
	}, err
}

// GetTaskFunc 根据任务ID获取任务函数
// GetTaskFunc 根据任务ID获取任务函数
func (c *Cron) GetTaskFunc(taskID int) (GetTask, error) {
	taskFunc, exists := c.TaskFuncMaps[taskID]
	if !exists {
		return GetTask{}, fmt.Errorf("task function not found for task ID: %d", taskID)
	}
	return taskFunc, nil
}

type GetTask struct {
	ID       int
	Func     func()
	FuncName string
}

func (c *Cron) GetTask() ([]GetTask, error) {
	var gt []GetTask
	for id, tf := range c.TaskFuncMaps {
		gt = append(gt, GetTask{ID: id, Func: tf.Func, FuncName: tf.FuncName})
	}
	return gt, nil
}

// ScheduleTasks1 根据配置调度任务
func (c *Cron) ScheduleTasks1(tasks []Task) error {
	if len(tasks) != 0 {
		for _, task := range tasks {
			taskFunc, err := c.GetTaskFunc(task.ID)
			if err != nil {
				return err
			}
			_, err = c.scheduler.NewJob(
				gocron.DurationJob(time.Duration(task.Interval)*time.Second),
				gocron.NewTask(taskFunc.Func),
			)
			if err != nil {
				return err
			}
		}
	}
	c.scheduler.Start()
	return nil
}

// 任务函数
func MyTask1() {
	fmt.Println("Task 1 is running at", time.Now())
}

func MyTask2() {
	fmt.Println("Task 2 is running at", time.Now())
}

func MyTask3() {
	fmt.Println("Task 3 is running at", time.Now())
}

func main() {
	cf := []Task{
		{
			ID:       1,
			Interval: 10,
		},
		{
			ID:       2,
			Interval: 15,
		},
		{
			ID:       3,
			Interval: 26,
		},
	}

	cron, err := NewCron()
	if err != nil {
		fmt.Println("Error creating Cron:", err)
		return
	}

	err = cron.ScheduleTasks1(cf)
	if err != nil {
		fmt.Println("Error scheduling tasks:", err)
		return
	}
	gt, err := cron.GetTask()
	if err != nil {
		fmt.Println("Error GetTask tasks:", err)
		return
	}
	for _, v := range gt {
		fmt.Println("tasks:", v.ID, v.FuncName)
	}

	select {} // Block forever
}
