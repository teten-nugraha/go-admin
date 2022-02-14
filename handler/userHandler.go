package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
	"golang.org/x/crypto/bcrypt"
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

	password, _ := bcrypt.GenerateFromPassword([]byte("1234"), 14)
	user.Password = string(password)

	db.DB.Create(&user)

	return c.JSON(user)
}
