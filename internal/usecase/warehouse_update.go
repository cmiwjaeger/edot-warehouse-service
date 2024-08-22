package usecase

import (
	"context"
	"edot-monorepo/services/warehouse-service/internal/entity"
	"edot-monorepo/services/warehouse-service/internal/gateway/messaging"
	"edot-monorepo/services/warehouse-service/internal/model"
	"edot-monorepo/services/warehouse-service/internal/model/converter"
	"time"

	"github.com/gofiber/fiber/v2"
)

type WarehouseUpdateUseCase struct {
	*WarehouseBaseUseCase
	WarehouseUpdateProducer *messaging.WarehouseProducer[model.Event]
}

func NewWarehouseUpdateUseCase(warehouseBaseUseCase *WarehouseBaseUseCase, warehouseUpdateProducer *messaging.WarehouseProducer[model.Event]) *WarehouseUpdateUseCase {
	return &WarehouseUpdateUseCase{
		WarehouseBaseUseCase:    warehouseBaseUseCase,
		WarehouseUpdateProducer: warehouseUpdateProducer,
	}
}

func (c *WarehouseUpdateUseCase) Exec(ctx context.Context, request *model.WarehouseUpdateRequest) (*model.WarehouseResponse, error) {

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	warehouse := &entity.Warehouse{
		ID:        request.ID,
		Name:      request.Name,
		Status:    true,
		UpdatedAt: time.Now(),
	}

	if err := c.WarehouseRepository.Update(tx, warehouse); err != nil {
		c.Log.Warnf("Failed create warehouse to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	event := converter.WarehouseToEventUpdated(warehouse)
	if err := c.WarehouseUpdateProducer.Send(event); err != nil {
		c.Log.WithError(err).Error("error publishing contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.WarehouseToResponse(warehouse), nil
}
