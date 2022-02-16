package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/middleware"
	"github.com/teten-nugraha/go-admin/model"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {

	if err := middleware.IsAuthorized(c, "users"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(model.Paginate(db.DB, &model.User{}, page))
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
func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := model.User{
		Id: uint(id),
	}

	db.DB.Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := model.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	db.DB.Model(&user).Updates(&user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := model.User{
		Id: uint(id),
	}

	db.DB.Delete(&user)

	return nil
}
