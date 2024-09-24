package config

import (
	route "go-clean-architecture/internal/delivery/http"
	"go-clean-architecture/internal/delivery/http/controller"
	"go-clean-architecture/internal/delivery/middleware"
	"go-clean-architecture/internal/repository"
	"go-clean-architecture/internal/usecase"

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
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	healthRepository := repository.NewUserRepository(config.Log)
	userRepository := repository.NewUserRepository(config.Log)

	// setup use cases
	healthUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, healthRepository)
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)

	// setup controller
	healthController := http.NewHealthController(healthUseCase, config.Log)
	userController := http.NewUserController(userUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:              config.App,
		UserController:   userController,
		HealthController: healthController,
		AuthMiddleware:   authMiddleware,
	}
	routeConfig.Setup()
}
