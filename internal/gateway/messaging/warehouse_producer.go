package messaging

import (
	"edot-monorepo/services/warehouse-service/internal/model"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type WarehouseProducer struct {
	Producer[*model.WarehouseCreatedEvent]
}

func NewWarehouseProducer(topic string, producer *kafka.Producer, log *logrus.Logger) *WarehouseProducer {
	return &WarehouseProducer{
		Producer: Producer[*model.WarehouseCreatedEvent]{
			Producer: producer,
			Topic:    topic,
			Log:      log,
		},
	}
}
