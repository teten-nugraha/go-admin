package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
)

func AllUsers(c *fiber.Ctx) error {

	var users []model.User

	db.DB.Find(&users)

	return c.JSON(&users)
}

func CreateUser(c *fiber.Ctx) error {

	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword(user.Password)

	db.DB.Create(&user)

	return c.JSON(user)
}
