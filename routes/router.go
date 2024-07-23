package routes

import (
	"github.com/gofiber/fiber/v2"
	"uaspw2/controller"
)

func SetupUserRoutes(app *fiber.App, controller controller.UserController) {
	userGroup := app.Group("/api/users")

	userGroup.Get("/", controller.FindAll)
	userGroup.Get("/:userId", controller.FindById)
	userGroup.Post("/", controller.Create)
	userGroup.Put("/:userId", controller.Update)
	userGroup.Delete("/:userId", controller.Delete)
}