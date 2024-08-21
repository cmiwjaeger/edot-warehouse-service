package controller

import (
	"edot-monorepo/services/warehouse-service/internal/model"
	"edot-monorepo/services/warehouse-service/internal/usecase"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WarehouseController struct {
	warehouseCreateUseCase *usecase.WarehouseCreateUseCase
	Log                    *logrus.Logger
	Validate               *validator.Validate
}

func NewWarehouseController(warehouseCreateUseCase *usecase.WarehouseCreateUseCase, log *logrus.Logger, validate *validator.Validate) *WarehouseController {
	return &WarehouseController{
		warehouseCreateUseCase: warehouseCreateUseCase,
		Log:                    log,
		Validate:               validate,
	}
}

func (c *WarehouseController) Create(ctx *fiber.Ctx) error {

	return ctx.JSON(model.WebResponse[*model.WarehouseResponse]{Data: nil})
}
