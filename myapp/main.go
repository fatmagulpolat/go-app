package main

import (
	"fmt"
	"myapp/internal/my-app-rest-api/application/controller"
	"myapp/internal/my-app-rest-api/application/handler/user"
	"myapp/internal/my-app-rest-api/application/pkg/server"
	"myapp/internal/my-app-rest-api/application/query"
	"myapp/internal/my-app-rest-api/application/repository"
	"myapp/internal/my-app-rest-api/application/web"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

func main() {
	app := fiber.New()

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
	/*


		app.Post("/user", func(ctx *fiber.Ctx) error {
			fmt.Println("hellö my first post endpoint")

			var request UserCreateRequest
			err := ctx.BodyParser(&request)

			if err != nil {
				fmt.Sprintf("error error %v ", err.Error())
				return err
			}
			if errors := Validate(request); errors != nil && errors[0].HasError {
				var errorResponse ErrorResponse
				var errorDetails []ErrorResponseDetail

				for _, validationError := range errors {
					var errorDetail ErrorResponseDetail
					errorDetail.FieldName = validationError.Field
					errorDetail.Description = fmt.Sprintf("%s filedhas an error", validationError.Field)
					errorDetails = append(errorDetails, errorDetail)
				}
				errorResponse.Status = http.StatusBadRequest
				errorResponse.ErrorDetail = errorDetails

				return ctx.Status(http.StatusBadRequest).JSON(errorResponse)
			}

			responseMessage := fmt.Sprintf("%s created succesfully", request.FirtsName)
			return ctx.Status(http.StatusOK).JSON(responseMessage)
		})
	*/
	//app.Listen(":3000")

}
