package active

import (
	"context"
	"encoding/json"
	"etcd/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (s *SystemCron) UpdateIDs2(task models.Tasks, jid uuid.UUID) error {
	key := fmt.Sprintf("%s%s", taskKeyPrefix, task.ID)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 增加超时时间
	defer cancel()

	task.Operation = models.OperationType(models.Nothing)
	value, err := json.Marshal(task)
	if err != nil {
		return fmt.Errorf("failed to marshal updated record: %v", err)
	}

	txn := Ed.Txn(ctx)
	_, err = txn.If(
		clientv3.Compare(clientv3.Version(key), "=", 0), // 如果 key 不存在
	).Then(
		clientv3.OpPut(key, string(value)),
	).Else(
		clientv3.OpPut(key, string(value)), // 如果已经存在，更新内容
	).Commit()
	if err != nil {
		return fmt.Errorf("failed to put updated record: %v", err)
	}
	return nil
}
