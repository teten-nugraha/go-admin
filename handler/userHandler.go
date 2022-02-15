package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
	"math"
	"strconv"
)

func AllUsers(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit
	var total int64

	var users []model.User

	db.DB.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	db.DB.Model(&model.User{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
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
