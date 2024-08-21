package messaging

import (
	"edot-monorepo/services/warehouse-service/internal/model"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type WarehouseConsumer struct {
	Log *logrus.Logger
}

func NewWarehouseConsumer(log *logrus.Logger) *WarehouseConsumer {
	return &WarehouseConsumer{
		Log: log,
	}
}

func (c WarehouseConsumer) Consume(message *kafka.Message) error {
	ContactEvent := new(model.Warehouse)
	if err := json.Unmarshal(message.Value, ContactEvent); err != nil {
		c.Log.WithError(err).Error("error unmarshalling Contact event")
		return err
	}

	// TODO process event
	c.Log.Infof("Received topic contacts with event: %v from partition %d", ContactEvent, message.TopicPartition.Partition)
	return nil
}
