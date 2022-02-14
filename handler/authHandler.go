package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/model"
)

func Register(c *fiber.Ctx) error {

	user := model.User{
		FirstName: "Teten",
	}

	user.LastName = "NUgraha"
	user.Email = "teten.nugraha18@gmail.com"
	user.Password = "Secret"

	return c.JSON(user)
}
