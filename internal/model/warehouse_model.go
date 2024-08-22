package model

import (
	"time"

	"github.com/google/uuid"
)

type WarehouseResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type WarehouseCreateRequest struct {
	Name string `json:"name"`
}

type WarehouseListRequest struct {
	QueryListRequest
}
