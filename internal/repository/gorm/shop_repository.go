package repository

import (
	"edot-monorepo/services/warehouse-service/internal/entity"

	"github.com/sirupsen/logrus"
)

type ShopRepository struct {
	Repository[entity.Warehouse]
	Log *logrus.Logger
}

func NewShopRepository(log *logrus.Logger) *ShopRepository {
	return &ShopRepository{
		Log: log,
	}
}
