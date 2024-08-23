package config

import (
	"edot-monorepo/services/warehouse-service/internal/delivery/http/controller"
	"edot-monorepo/services/warehouse-service/internal/delivery/http/route"
	"edot-monorepo/services/warehouse-service/internal/gateway/messaging"
	repository "edot-monorepo/services/warehouse-service/internal/repository/gorm"
	"edot-monorepo/services/warehouse-service/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
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
	// Producer *kafka.Producer
	Reader *kafka.Reader
	Writer *kafka.Writer
}

func Bootstrap(config *BootstrapConfig) {

	warehouseRepository := repository.NewWarehouseRepository(config.Log)

	producer := messaging.NewProducer(config.Writer, config.Log)

	warehouseBaseUseCase := usecase.NewWarehouseUseCase(config.DB, config.Log, warehouseRepository, config.Validate, producer)
	warehouseCreateUseCase := usecase.NewWarehouseCreateUseCase(warehouseBaseUseCase)
	warehouseUpdateUseCase := usecase.NewWarehouseUpdateUseCase(warehouseBaseUseCase)
	warehouseListUseCase := usecase.NewWarehouseListUseCase(warehouseBaseUseCase)

	warehouseController := controller.NewWarehouseController(warehouseCreateUseCase, warehouseUpdateUseCase, warehouseListUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:                 config.App,
		WarehouseController: warehouseController,
	}

	routeConfig.Setup()
}
