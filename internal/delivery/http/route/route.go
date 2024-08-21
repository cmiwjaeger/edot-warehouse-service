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
func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/create", c.WarehouseController.Create)

}

func (c *RouteConfig) SetupAuthRoute() {
	// c.App.Use(c.AuthMiddleware)
}
