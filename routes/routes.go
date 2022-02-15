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

	app.Get("/api/users/info", handler.UpdateInfo)
	app.Get("/api/users/password", handler.UpdatePassword)

	app.Get("/api/users", handler.AllUsers)
	app.Post("/api/users", handler.CreateUser)
	app.Put("/api/users/:id", handler.UpdateUser)
	app.Get("/api/users/:id", handler.GetUser)
	app.Delete("/api/users/:id", handler.DeleteUser)

	app.Get("/api/roles", handler.AllRoles)
	app.Post("/api/roles", handler.CreateRole)
	app.Put("/api/roles/:id", handler.UpdateRole)
	app.Get("/api/roles/:id", handler.GetRole)
	app.Delete("/api/roles/:id", handler.DeleteRole)

	app.Get("/api/permissions", handler.GetAllPermission)
}
