package active

import (
	"context"
	"encoding/json"
	"errors"
	"etcd/models"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	taskKeyPrefix = "/system/cron/tasks/"
)

var Api = new(SystemCron)

type SystemCron struct{}

func (s SystemCron) Tables(page, pageSize int, name string) (*models.TaskCounts, error) {
	if Ed == nil {
		return nil, fmt.Errorf("etcdClient is not initialized")
	}

	var data []models.Tasks
	start := (page - 1) * pageSize
	end := start + pageSize

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := Ed.Get(ctx, taskKeyPrefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))

	if err != nil {
		return nil, err
	}

	for _, kv := range resp.Kvs {
		var record models.Tasks
		if err := json.Unmarshal(kv.Value, &record); err != nil {
			return nil, err
		}
		// 如果 name 不为空且匹配，则添加到 records 中并立即返回
		if name != "" && record.Name == name {
			return &models.TaskCounts{
				Items: []models.Tasks{record},
				Total: len([]models.Tasks{record}),
			}, nil
		}
		// 如果 name 为空，则添加到 data 中
		if name == "" {
			data = append(data, record)
		}
	}

	total := len(data)
	if start >= total {
		return &models.TaskCounts{
			Items: data,
			Total: total,
		}, nil
	}
	if end > total {
		end = total
	}
	return &models.TaskCounts{
		Items: data,
		Total: total,
	}, nil
}
func (s *SystemCron) Adds(data *models.Tasks) error {
	if err := s.check(data.Name); err != nil {
		return err
	}
	id := uuid.New().String()
	key := fmt.Sprintf("%s%s", taskKeyPrefix, id)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	record := []models.Tasks{models.Tasks{
		ID:          id,
		Name:        data.Name,
		Status:      data.Status,
		Mark:        data.Mark,
		Func:        data.Func,
		FuncID:      data.FuncID,
		Interval:    data.Interval,
		CreatedTime: NowTime(),
		Operation:   models.OperationType(models.Add),
	}}

	value, err := json.Marshal(record)
	if err != nil {
		return err
	}
	_, err = Ed.Put(ctx, key, string(value))
	if err != nil {
		return err
	}
	return nil
}

func (s *SystemCron) Updates(data *models.Tasks) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	key := fmt.Sprintf("%s%s", taskKeyPrefix, data.ID)

	data.UpdatedTime = NowTime()
	data.Operation = models.OperationType(models.Update)
	// 如果id不存在
	if !s.GetTask(ctx, key) {
		return errors.New("非法参数")
	}

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = Ed.Put(ctx, key, string(value))
	if err != nil {
		return err
	}

	return nil
}
func (*SystemCron) Dels(id string) error {
	key := fmt.Sprintf("%s%s", taskKeyPrefix, id)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := Ed.Delete(ctx, key)
	return err
}

func (SystemCron) GetTask(ctx context.Context, key string) bool {
	resp, err := Ed.Get(ctx, key)
	if err != nil {
		return false
	}
	if len(resp.Kvs) != 0 {
		return true
	}
	return false
}

func (s *SystemCron) check(name string) error {
	if Ed == nil {
		return errors.New("初始化失败")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := Ed.Get(ctx, taskKeyPrefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return err
	}

	for _, kv := range resp.Kvs {
		var record models.Tasks
		if err := json.Unmarshal(kv.Value, &record); err != nil {
			return err
		}
		if record.Name == name {
			return errors.New("重复存在")
		}
	}
	return nil
}

func getSplit(key string) string {
	parts := strings.Split(key, "/")
	if len(parts) > 0 {
		uuid := parts[len(parts)-1]
		fmt.Println("UUID:", uuid)
		return uuid
	} else {
		fmt.Println("Invalid key format")
	}
	return ""
}
