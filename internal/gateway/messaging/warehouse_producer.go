package messaging

import (
	"edot-monorepo/services/warehouse-service/internal/model"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type WarehouseProducer[T model.Event] struct {
	Producer[T]
}

func NewWarehouseProducer[T model.Event](topic string, producer *kafka.Producer, log *logrus.Logger) *WarehouseProducer[model.Event] {

	return &WarehouseProducer[model.Event]{
		Producer: Producer[model.Event]{
			Producer: producer,
			Topic:    topic,
			Log:      log,
		},
	}
}
