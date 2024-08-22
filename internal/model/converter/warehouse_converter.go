package converter

import (
	"edot-monorepo/services/warehouse-service/internal/entity"
	"edot-monorepo/services/warehouse-service/internal/model"
)

func WarehouseToResponse(item *entity.Warehouse) *model.WarehouseResponse {
	return &model.WarehouseResponse{
		ID:        item.ID,
		Name:      item.Name,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func WarehouseToEvent(item *entity.Warehouse) *model.WarehouseCreatedEvent {
	return &model.WarehouseCreatedEvent{
		ID:        item.ID,
		Name:      item.Name,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func WarehouseListToResponseList(items []entity.Warehouse) []*model.WarehouseResponse {
	productResponse := make([]*model.WarehouseResponse, len(items))

	for i, item := range items {

		productResponse[i] = WarehouseToResponse(&item)
	}

	return productResponse
}
