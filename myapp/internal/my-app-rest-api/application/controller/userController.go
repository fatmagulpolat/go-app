package controller

import (
	"fmt"
	"myapp/internal/my-app-rest-api/application/controller/request"
	"myapp/internal/my-app-rest-api/application/controller/response"
	"myapp/internal/my-app-rest-api/application/handler/user"
	"myapp/internal/my-app-rest-api/application/query"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	GetUserById(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
}

type userController struct {
	userQueryService   query.IUserQueryService
	userCommandHandler user.ICommandHandler
}

func NewUserController(userQueryService query.IUserQueryService,
	userCommandHandler user.ICommandHandler) IUserController {
	return &userController{
		userQueryService:   userQueryService,
		userCommandHandler: userCommandHandler,
	}
}

// GetById implements IUserController.
func (c *userController) GetUserById(ctx *fiber.Ctx) error {

	userId := ctx.Params("userId")
	if userId == "" {
		return ctx.Status(http.StatusBadRequest).JSON("userId can not be empty")
	}

	user, err := c.userQueryService.GetById(ctx.UserContext(), userId)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var userDto = response.ToUserResponse(user)
	return ctx.Status(http.StatusOK).JSON(userDto)
}

// GetUser implements IUserController.
func (c *userController) GetUser(ctx *fiber.Ctx) error {

	users, err := c.userQueryService.Get(ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	userDtoList := response.ToUserResponseList(users)

	return ctx.Status(http.StatusOK).JSON(userDtoList)
}

func (c *userController) Save(ctx *fiber.Ctx) error {

	var request request.UserCreateRequest
	err := ctx.BodyParser(&request)

	if err != nil {
		fmt.Sprintf("error error %v ", err.Error())
		return err
	}

	if err = c.userCommandHandler.Save(ctx.UserContext(), request.ToCommand()); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(http.StatusOK).JSON("user created successfully")
}
