package usecase

import (
	"context"
	"go-clean-architecture/internal/model"
	"go-clean-architecture/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (c *UserUseCase) Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	return nil, fiber.ErrInternalServerError
}

func (c *UserUseCase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error) {
	return nil, fiber.ErrInternalServerError
}

func (c *UserUseCase) Logout(ctx context.Context, request *model.LogoutUserRequest) (bool, error) {
	return true, nil
}

func (c *UserUseCase) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	return nil, fiber.ErrBadRequest
}

func (c *UserUseCase) Current(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error) {
	return nil, fiber.ErrBadRequest
}

func (c *UserUseCase) Update(ctx context.Context, request *model.UpdateUserRequest) (*model.UserResponse, error) {
	return nil, fiber.ErrBadRequest
}
