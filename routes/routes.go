package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/handler"
	"github.com/teten-nugraha/go-admin/middleware"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/register", handler.Register)
	app.Post("/api/login", handler.Login)

	app.Use(middleware.IsAuthenticated)

	app.Post("/api/logout", handler.Logout)
	app.Get("/api/user", handler.User)

	app.Get("/api/users", handler.AllUsers)
	app.Post("/api/users", handler.CreateUser)
	app.Put("/api/users/:id", handler.UpdateUser)
	app.Get("/api/users/:id", handler.GetUser)
	app.Delete("/api/users/:id", handler.DeleteUser)
}
