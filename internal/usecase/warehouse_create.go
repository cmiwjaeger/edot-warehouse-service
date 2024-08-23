package usecase

import (
	"context"
	"edot-monorepo/services/warehouse-service/internal/entity"
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

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	warehouse := &entity.Warehouse{
		Name:   request.Name,
		Status: true,
	}

	if err := c.WarehouseRepository.Create(tx, warehouse); err != nil {
		c.Log.Warnf("Failed create warehouse to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	event := converter.WarehouseToEvent(warehouse)
	if err = c.Producer.Produce(ctx, "shop_created", event); err != nil {
		c.Log.WithError(err).Error("error publishing contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.WarehouseToResponse(warehouse), nil
}
