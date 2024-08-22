package usecase

import (
	"context"
	"edot-monorepo/services/warehouse-service/internal/entity"
	"edot-monorepo/services/warehouse-service/internal/model"
	"edot-monorepo/services/warehouse-service/internal/model/converter"

	"github.com/gofiber/fiber/v2/log"
)

type WarehouseListUseCase struct {
	*WarehouseBaseUseCase
}

func NewWarehouseListUseCase(warehouseBaseUseCase *WarehouseBaseUseCase) *WarehouseListUseCase {

	return &WarehouseListUseCase{
		warehouseBaseUseCase,
	}
}

func (u *WarehouseListUseCase) Exec(ctx context.Context, request *model.WarehouseListRequest) ([]*model.WarehouseResponse, error) {
	warehouses := make([]entity.Warehouse, 0)

	err := u.WarehouseRepository.FindAll(u.DB, &warehouses)
	if err != nil {
		log.Errorf("%v", err)
	}

	return converter.WarehouseListToResponseList(warehouses), nil
}
