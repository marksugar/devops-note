package models

type Test struct {
	ID     int
	Name   string
	Status int
}
type TaskMessage struct {
	Operation OperationType
	Data      []Test
}
type OperationType string

const (
	Add    OperationType = "add"
	Update OperationType = "update"
	Delete OperationType = "delete"
)
