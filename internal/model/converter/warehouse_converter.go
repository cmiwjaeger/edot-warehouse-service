package converter

import (
	"edot-monorepo/services/warehouse-service/internal/entity"
	"edot-monorepo/services/warehouse-service/internal/model"
)

func WarehouseToResponse(user *entity.Warehouse) *model.WarehouseResponse {
	return &model.WarehouseResponse{
		ID: user.ID,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func WarehouseToTokenResponse(user *entity.Warehouse) *model.WarehouseResponse {
	return &model.WarehouseResponse{}
}
