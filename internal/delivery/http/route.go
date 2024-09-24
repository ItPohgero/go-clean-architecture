package route

import (
	http "go-clean-architecture/internal/delivery/http/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App              *fiber.App
	UserController   *http.UserController
	HealthController *http.HealthController
	AuthMiddleware   fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/health", c.HealthController.Health)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/api/users", c.UserController.Current)
}
