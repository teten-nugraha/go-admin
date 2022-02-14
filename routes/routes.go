package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.Hello)
}
