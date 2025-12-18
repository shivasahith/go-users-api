package routes

import (
	"go-users-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
	app.Get("/users", userHandler.ListUsers)
}
