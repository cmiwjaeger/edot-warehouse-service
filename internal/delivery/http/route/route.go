package route

import (
	http "edot-monorepo/services/warehouse-service/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App

	WarehouseController *http.WarehouseController
	AuthMiddleware      fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

const wareHouseRoute string = "/api/warehouse"

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get(wareHouseRoute, c.WarehouseController.List)
	c.App.Post(wareHouseRoute, c.WarehouseController.Create)
	c.App.Put(wareHouseRoute, c.WarehouseController.Update)

}

func (c *RouteConfig) SetupAuthRoute() {
	// c.App.Use(c.AuthMiddleware)

}
