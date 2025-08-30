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

// GetUserById godoc
// @Summary      This method get user by id.
// @Description  get user by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        userId   path      string  true  "userId"
// @Success      200  {object}  response.UserResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/user/{userId} [get]
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

// GetUser godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {array}   response.UserResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/user [get]
func (c *userController) GetUser(ctx *fiber.Ctx) error {

	users, err := c.userQueryService.Get(ctx.UserContext())
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	userDtoList := response.ToUserResponseList(users)

	return ctx.Status(http.StatusOK).JSON(userDtoList)
}

// SaveUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the given information
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body      request.UserCreateRequest  true  "User create request"
// @Success      200      {string}  string  "user created successfully"
// @Failure      400      {string}  string  "Bad request"
// @Failure      500      {string}  string  "Internal server error"
// @Router       /api/v1/user [post]
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
