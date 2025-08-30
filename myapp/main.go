package main

import (
	"fmt"
	"myapp/configuration"
	"myapp/internal/my-app-rest-api/application/controller"
	"myapp/internal/my-app-rest-api/application/handler/user"
	"myapp/internal/my-app-rest-api/application/pkg/server"
	"myapp/internal/my-app-rest-api/application/query"
	"myapp/internal/my-app-rest-api/application/repository"
	"myapp/internal/my-app-rest-api/application/web"
	"net/http"

	_ "myapp/docs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/google/uuid"
)

type CustomValidationError struct {
	HasError bool
	Field    string
	Tag      string
	Param    string
	Value    interface{}
}

type ErrorResponse struct {
	Status      int32                 `json:"status"`
	ErrorDetail []ErrorResponseDetail `json:"errorDetail"`
}

type ErrorResponseDetail struct {
	FieldName   string `json:"fieldName"`
	Description string `json:"description"`
}

var validate = validator.New()

func Validate(data interface{}) []CustomValidationError {

	var customerValidationError []CustomValidationError
	if errors := validate.Struct(data); errors != nil {
		for _, fieldError := range errors.(validator.ValidationErrors) {
			var cve CustomValidationError
			cve.HasError = true
			cve.Field = fieldError.Error()
			cve.Tag = fieldError.Tag()
			cve.Param = fieldError.Param()
			cve.Value = fieldError.Value()
			customerValidationError = append(customerValidationError, cve)
		}
	}
	return customerValidationError
}

// @title go app
// @version 1.0
// @description This is a sample swagger for Fiber
// @contact.name API Support
// @license.name Apache 2.0
func main() {
	app := fiber.New()

	configureSwaggerUi(app)

	userRepository := repository.NewUserRepository()
	userQueryService := query.NewUserQueryService(userRepository)
	userCommandHandler := user.NewCommandHandler(userRepository)

	userController := controller.NewUserController(userQueryService, userCommandHandler)

	web.InitRouter(app, userController)

	server.NewServer(app).StartHttpServer()

	app.Use(func(ctx *fiber.Ctx) error {
		fmt.Println("Hello client you are call my %s%s and method %s", ctx.BaseURL(), ctx.Request().RequestURI(), ctx.Request().Header.Method())
		return ctx.Next()
	})

	app.Use("/user", func(ctx *fiber.Ctx) error {
		correlationId := ctx.Get("X-CorrelationId")

		if correlationId == "" {
			return ctx.Status(http.StatusBadRequest).JSON("You have to send correlationid")
		}

		_, err := uuid.Parse(correlationId)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON("CorrelationId must be guid")
		}

		ctx.Locals("correlationId", correlationId)
		return ctx.Next()
	})

	//custom validation tag
	validate.RegisterValidation("acceptAge", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() > 18
	})

}

func configureSwaggerUi(app *fiber.App) {
	if configuration.Env != "prod" {
		app.Get("/swagger/*", swagger.HandlerDefault)

		app.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusMovedPermanently).Redirect("/swagger/index.html")
		})
	}
}
