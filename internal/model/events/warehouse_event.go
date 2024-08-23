package events

import (
	"time"

	"github.com/google/uuid"
)

type WarehouseCreatedEvent struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (c *WarehouseCreatedEvent) GetId() string {
	return c.ID.String()
}

type WarehouseUpdatedEvent struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (c *WarehouseUpdatedEvent) GetId() string {
	return c.ID.String()
}
