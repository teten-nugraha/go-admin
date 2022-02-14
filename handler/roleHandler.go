package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
	"strconv"
)

func AllRoles(c *fiber.Ctx) error {

	var roles []model.Role

	db.DB.Find(&roles)

	return c.JSON(&roles)
}

func CreateRole(c *fiber.Ctx) error {

	var role model.Role
	if err := c.BodyParser(&role); err != nil {
		return err
	}

	db.DB.Create(&role)

	return c.JSON(role)
}
func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := model.Role{
		Id: uint(id),
	}

	db.DB.Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := model.Role{
		Id: uint(id),
	}

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	db.DB.Model(&role).Updates(&role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := model.Role{
		Id: uint(id),
	}

	db.DB.Delete(&role)

	return nil
}
