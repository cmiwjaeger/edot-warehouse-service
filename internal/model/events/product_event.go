package events

import "github.com/google/uuid"

type ProductCreatedEvent struct {
	ID    uuid.UUID        `json:"uuid"`
	Name  string           `json:"name"`
	Stock int64            `json:"stock"`
	Shop  ShopCreatedEvent `json:"shop"`
}

func (c *ProductCreatedEvent) GetId() string {
	return c.ID.String()
}
