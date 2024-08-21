package usecase

import (
	repository "edot-monorepo/services/warehouse-service/internal/repository/gorm"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WarehouseBaseUseCase struct {
	DB                  *gorm.DB
	Log                 *logrus.Logger
	WarehouseRepository *repository.WarehouseRepository
	Validate            *validator.Validate
}

func NewWarehouseUseCase(db *gorm.DB, log *logrus.Logger, warehouseRepo *repository.WarehouseRepository, validate *validator.Validate) *WarehouseBaseUseCase {
	return &WarehouseBaseUseCase{
		DB:                  db,
		Log:                 log,
		WarehouseRepository: warehouseRepo,
		Validate:            validate,
	}
}
