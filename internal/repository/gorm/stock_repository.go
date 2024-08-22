package repository

import (
	"edot-monorepo/services/warehouse-service/internal/entity"

	"github.com/sirupsen/logrus"
)

type StockRepository struct {
	Repository[entity.Warehouse]
	Log *logrus.Logger
}

func NewStockRepository(log *logrus.Logger) *StockRepository {
	return &StockRepository{
		Log: log,
	}
}
