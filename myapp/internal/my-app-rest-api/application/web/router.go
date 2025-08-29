package web

import (
	"myapp/internal/my-app-rest-api/application/controller"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, userContoller controller.IUserController) {
	app.Get("/healthcheck", func(context *fiber.Ctx) error { return context.SendStatus(http.StatusOK) })

	routeGroup := app.Group("/api/v1")
	routeGroup.Get("/user", userContoller.GetUser)
	routeGroup.Get("/user/:userId", userContoller.GetUserById)
	routeGroup.Post("/user", userContoller.Save)
}
