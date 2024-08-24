package messaging

import (
	"edot-monorepo/services/warehouse-service/internal/entity"
	"edot-monorepo/shared/events"

	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WarehouseConsumer struct {
	Log       *logrus.Logger
	DB        *gorm.DB
	Validator *validator.Validate
}

func NewWarehouseConsumer(log *logrus.Logger, db *gorm.DB, validate *validator.Validate) *WarehouseConsumer {
	return &WarehouseConsumer{
		Log:       log,
		DB:        db,
		Validator: validate,
	}
}

func (c WarehouseConsumer) ConsumeShopCreated(message *kafka.Message) error {
	event := new(events.ShopCreatedEvent)
	if err := json.Unmarshal(message.Value, event); err != nil {
		c.Log.WithError(err).Error("error unmarshalling ShopCreatedEvent event")
		return err
	}
	data := &entity.Shop{
		Name:    event.Name,
		Address: event.Address,
	}

	err := c.DB.Create(data).Error
	if err != nil {
		c.Log.WithError(err).Error("error insert into db")
	}

	c.Log.Infof("Received topic  with event: %v from partition %s", event, message.Topic)
	return nil
}

func (c WarehouseConsumer) ConsumeStockChanged(message *kafka.Message) error {
	event := new(events.WarehouseCreatedEvent)
	if err := json.Unmarshal(message.Value, event); err != nil {
		c.Log.WithError(err).Error("error unmarshalling WarehouseCreatedEvent event")
		return err
	}

	// TODO process event
	c.Log.Infof("Received topic contacts with event: %v from partition %s", event, message.Topic)
	return nil
}

func (c WarehouseConsumer) ConsumeShopWarehouseAssigned(message *kafka.Message) (err error) {
	event := new(events.ShopWarehouseAssignedEvent)
	if err := json.Unmarshal(message.Value, event); err != nil {
		c.Log.WithError(err).Error("error unmarshalling WarehouseCreatedEvent event")
		return err
	}

	data := &entity.WarehouseShop{
		WarehouseID: event.WarehouseID,
		ShopID:      event.ShopID,
	}
	if event.Assigned {
		err = c.DB.Create(data).Error
	} else {
		err = c.DB.Delete(data).Error
	}
	if err != nil {
		c.Log.WithError(err).Error("error insert into db")
	}

	c.Log.Infof("Received topic contacts with event: %v from partition %s", event, message.Topic)
	return nil
}
