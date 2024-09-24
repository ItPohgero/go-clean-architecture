package http

import (
	"go-clean-architecture/internal/model"
	"go-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewHealthController(useCase *usecase.UserUseCase, logger *logrus.Logger) *HealthController {
	return &HealthController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *HealthController) Health(ctx *fiber.Ctx) error {
	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
