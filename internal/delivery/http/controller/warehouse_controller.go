package controller

import (
	"edot-monorepo/services/warehouse-service/internal/model"
	"edot-monorepo/services/warehouse-service/internal/usecase"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WarehouseController struct {
	warehouseCreateUseCase *usecase.WarehouseCreateUseCase
	warehouseListUseCase   *usecase.WarehouseListUseCase
	Log                    *logrus.Logger
	Validate               *validator.Validate
}

func NewWarehouseController(warehouseCreateUseCase *usecase.WarehouseCreateUseCase, warehouseListUseCase *usecase.WarehouseListUseCase, log *logrus.Logger, validate *validator.Validate) *WarehouseController {
	return &WarehouseController{
		warehouseCreateUseCase: warehouseCreateUseCase,
		warehouseListUseCase:   warehouseListUseCase,
		Log:                    log,
		Validate:               validate,
	}
}

func (c *WarehouseController) Create(ctx *fiber.Ctx) error {

	request := new(model.WarehouseCreateRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.warehouseCreateUseCase.Exec(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.WarehouseResponse]{Data: response})
}

func (c *WarehouseController) List(ctx *fiber.Ctx) error {
	request, err := parseQueryToModel(ctx)
	if err != nil {
		c.Log.Warnf("Failed to parse query : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.warehouseListUseCase.Exec(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to list products : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]*model.WarehouseResponse]{
		Data: response,
	})

}

func parseQueryToModel(ctx *fiber.Ctx) (*model.WarehouseListRequest, error) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		return nil, err
	}

	size, err := strconv.Atoi(ctx.Query("size"))
	if err != nil {
		return nil, err
	}

	return &model.WarehouseListRequest{
		QueryListRequest: model.QueryListRequest{
			Keyword: ctx.Query("keyword"),
			Page:    page,
			Size:    size,
		},
	}, nil
}
