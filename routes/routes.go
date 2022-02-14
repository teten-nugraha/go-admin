package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/register", handler.Register)
	app.Post("/api/login", handler.Login)
}
