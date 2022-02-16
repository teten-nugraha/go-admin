package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
	"strconv"
)

func AllOrders(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(model.Paginate(db.DB, &model.Order{}, page))
}
