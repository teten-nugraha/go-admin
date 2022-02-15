package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
)

func GetAllPermission(c *fiber.Ctx) error {

	var permissions []model.Permission

	db.DB.Find(&permissions)

	return c.JSON(&permissions)
}
