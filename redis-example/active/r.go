package active

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"redis/models"

	"github.com/redis/go-redis/v9"
)

type SystemTest struct{}

func TaskChannel2(msgChan chan<- models.TaskMessage) {
	pubsub := Rdb.Subscribe(context.Background(), "taskchannel")
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		var receivedMessage models.TaskMessage
		if err := json.Unmarshal([]byte(msg.Payload), &receivedMessage); err != nil {
			log.Printf("Failed to unmarshal data: %v", err)
			continue
		}
		msgChan <- receivedMessage
	}
}
func (SystemTest) PublishTaskChange(rdb *redis.Client, operation models.OperationType, test models.Test) error {

	message := models.TaskMessage{
		Operation: operation,
		Data:      []models.Test{test},
	}

	messageBytes, err := json.Marshal(message)

	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	return rdb.Publish(ctx, "taskchannel", messageBytes).Err()
}
