package events

import (
	"time"

	"github.com/google/uuid"
)

type ShopCreatedEvent struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (c *ShopCreatedEvent) GetId() string {
	return c.ID.String()
}

type ShopWarehouseAssignedEvent struct {
	ID          uuid.UUID `json:"id"`
	Assigned    bool      `json:"assigned"`
	ShopID      uuid.UUID `json:"shop_id"`
	WarehouseID uuid.UUID `json:"warehouse_id"`
}

func (c *ShopWarehouseAssignedEvent) GetId() string {
	return c.ID.String()
}
