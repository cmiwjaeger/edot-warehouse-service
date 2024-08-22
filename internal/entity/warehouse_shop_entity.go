package entity

import "github.com/google/uuid"

type WarehouseShop struct {
	ID          uuid.UUID `gorm:"column:id"`
	ShopID      uuid.UUID `gorm:"column:shop_id"`
	WarehouseID uuid.UUID `gorm:"column:warehouse_id"`
}

func (u *WarehouseShop) TableName() string {
	return "warehouse_shops"
}
