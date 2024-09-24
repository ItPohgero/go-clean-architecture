package http

import (
	"go-clean-architecture/internal/delivery/middleware"
	"go-clean-architecture/internal/model"
	"go-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

// Register endpoint to register a new user
//
// @Summary Register a new user
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param data body model.RegisterUserRequest true "Register user data"
// @Success 201 {object} model.WebResponse{data=model.UserResponse}
// @Failure 400 {object} model.WebResponse{data=string}
// @Router /users/register [post]
func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)

	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

// Login endpoint to login a user
//
// @Summary Login a user
// @Description Login a user
// @Tags User
// @Accept json
// @Produce json
// @Param data body model.LoginUserRequest true "Login user data"
// @Success 200 {object} model.WebResponse{data=model.UserResponse}
// @Failure 400 {object} model.WebResponse{data=string}
// @Router /users/login [post]
func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

// Current endpoint to get current user
//
// @Summary Get current user
// @Description Get current user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=model.UserResponse}
// @Failure 400 {object} model.WebResponse{data=string}
// @Security BearerAuth
// @Router /users/current [get]
func (c *UserController) Current(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetUserRequest{
		ID: auth.ID,
	}

	response, err := c.UseCase.Current(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get current user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

// Logout endpoint to logout a user
//
// @Summary Logout a user
// @Description Logout a user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=bool}
// @Failure 400 {object} model.WebResponse{data=string}
// @Security BearerAuth
// @Router /users/logout [get]
func (c *UserController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.LogoutUserRequest{
		ID: auth.ID,
	}

	response, err := c.UseCase.Logout(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to logout user")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}

// Update endpoint to update a user
//
// @Summary Update a user
// @Description Update a user
// @Tags User
// @Accept json
// @Produce json
// @Param data body model.UpdateUserRequest true "Update user data"
// @Success 200 {object} model.WebResponse{data=model.UserResponse}
// @Failure 400 {object} model.WebResponse{data=string}
// @Security BearerAuth
// @Router /users/update [post]
func (c *UserController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.ID = auth.ID
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}
