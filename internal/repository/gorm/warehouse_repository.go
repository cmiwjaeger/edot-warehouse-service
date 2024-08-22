package repository

import (
	"edot-monorepo/services/warehouse-service/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WarehouseRepository struct {
	Repository[entity.Warehouse]
	Log *logrus.Logger
}

func NewWarehouseRepository(log *logrus.Logger) *WarehouseRepository {
	return &WarehouseRepository{
		Log: log,
	}
}

func (r *WarehouseRepository) FindAll(db *gorm.DB, entity *[]entity.Warehouse) error {
	return db.Find(entity).Error
}
