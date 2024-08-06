package models

import "github.com/google/uuid"

type Tasks struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Interval    int       `json:"interval"` // 以秒为单位
	Func        string    `json:"func"`     // 任务中函数名称
	FuncID      int       `json:"funcid"`   // 对应函数的id
	UUID        uuid.UUID `json:"uuid"`
	Status      int       `json:"status"`
	Mark        string    `json:"mark"`
	Operation   OperationType
	CreatedTime string `json:"create_time"`
	UpdatedTime string `json:"update_time"`
}
type Tablist struct {
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Name     string `json:"name"`
	Total    int    `json:"total"`
}
type TaskCounts struct {
	Items []Tasks `json:"items"`
	Total int     `json:"total"`
}
type OperationType string

const (
	Add     OperationType = "add"
	Update  OperationType = "update"
	Delete  OperationType = "delete"
	Nothing OperationType = "nothing"
)

type TaskMessage struct {
	Operation OperationType
	Data      []Tasks
}
