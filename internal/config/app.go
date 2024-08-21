package config

import (
	"edot-monorepo/services/warehouse-service/internal/delivery/http/controller"
	"edot-monorepo/services/warehouse-service/internal/delivery/http/route"
	repository "edot-monorepo/services/warehouse-service/internal/repository/gorm"
	"edot-monorepo/services/warehouse-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {

	warehouseRepository := repository.NewWarehouseRepository(config.Log)
	warehouseBaseUseCase := usecase.NewWarehouseUseCase(config.DB, config.Log, warehouseRepository, config.Validate)
	warehouseCreateUseCase := usecase.NewWarehouseCreateUseCase(warehouseBaseUseCase)

	warehouseController := controller.NewWarehouseController(warehouseCreateUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:                 config.App,
		WarehouseController: warehouseController,
	}

	routeConfig.Setup()
}
