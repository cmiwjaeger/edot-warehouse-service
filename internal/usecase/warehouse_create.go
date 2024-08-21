package usecase

import (
	"context"
	"edot-monorepo/services/warehouse-service/internal/model"
	"edot-monorepo/services/warehouse-service/internal/model/converter"

	"github.com/gofiber/fiber/v2"
)

type WarehouseCreateUseCase struct {
	*WarehouseBaseUseCase
}

func NewWarehouseCreateUseCase(warehouseBaseUseCase *WarehouseBaseUseCase) *WarehouseCreateUseCase {
	return &WarehouseCreateUseCase{
		warehouseBaseUseCase,
	}
}

func (c *WarehouseCreateUseCase) Exec(ctx context.Context, request *model.WarehouseCreateRequest) (*model.WarehouseResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.WarehouseRepository.Create(tx, nil); err != nil {
		c.Log.Warnf("Failed create warehouse to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.WarehouseToResponse(nil), nil
}
