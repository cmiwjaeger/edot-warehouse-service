package model

import (
	"time"

	"github.com/google/uuid"
)

type WarehouseResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type WarehouseCreateRequest struct {
	Name string `json:"name"`
}

type WarehouseUpdateRequest struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	Name      string    `json:"name"`
	Status    bool      `json:"status" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type WarehouseListRequest struct {
	QueryListRequest
}
