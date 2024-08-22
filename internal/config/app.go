package config

import (
	"edot-monorepo/services/warehouse-service/internal/delivery/http/controller"
	"edot-monorepo/services/warehouse-service/internal/delivery/http/route"
	"edot-monorepo/services/warehouse-service/internal/gateway/messaging"
	repository "edot-monorepo/services/warehouse-service/internal/repository/gorm"
	"edot-monorepo/services/warehouse-service/internal/usecase"
	"edot-monorepo/shared/events"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
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
	Producer *kafka.Producer
}

func Bootstrap(config *BootstrapConfig) {

	warehouseRepository := repository.NewWarehouseRepository(config.Log)

	warehouseCreatedProducer := messaging.NewWarehouseProducer[*events.WarehouseCreatedEvent]("warehouse_created", config.Producer, config.Log)
	warehouseStatusUpdatedProducer := messaging.NewWarehouseProducer[*events.WarehouseUpdatedEvent]("warehouse_updated", config.Producer, config.Log)

	warehouseBaseUseCase := usecase.NewWarehouseUseCase(config.DB, config.Log, warehouseRepository, config.Validate)
	warehouseCreateUseCase := usecase.NewWarehouseCreateUseCase(warehouseBaseUseCase, warehouseCreatedProducer)
	warehouseUpdateUseCase := usecase.NewWarehouseUpdateUseCase(warehouseBaseUseCase, warehouseStatusUpdatedProducer)
	warehouseListUseCase := usecase.NewWarehouseListUseCase(warehouseBaseUseCase)

	warehouseController := controller.NewWarehouseController(warehouseCreateUseCase, warehouseUpdateUseCase, warehouseListUseCase, config.Log, config.Validate)

	routeConfig := route.RouteConfig{
		App:                 config.App,
		WarehouseController: warehouseController,
	}

	routeConfig.Setup()
}
