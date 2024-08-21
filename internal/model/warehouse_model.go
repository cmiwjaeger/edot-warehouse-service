package model

import (
	"time"

	"github.com/google/uuid"
)

type Warehouse struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WarehouseResponse struct {
	ID uuid.UUID `json:"id,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type WarehouseCreateRequest struct {
	ID uuid.UUID
}
