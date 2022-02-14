package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {

	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "password doesnt match",
		})
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := model.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  string(encryptedPassword),
	}

	db.DB.Create(&user)

	return ctx.JSON(user)
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	var user model.User

	db.DB.Where("email = ?", data["email"]).Find(&user)

	if user.Id == 0 {
		ctx.Status(404)
		return ctx.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		ctx.Status(404)
		return ctx.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	return ctx.JSON(user)
}
