package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
	"strconv"
)

func AllProducts(c *fiber.Ctx) error {

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(model.Paginate(db.DB, &model.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {

	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return err
	}

	db.DB.Create(&product)

	return c.JSON(product)
}
func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := model.Product{
		Id: uint(id),
	}

	db.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := model.Product{
		Id: uint(id),
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	db.DB.Model(&product).Updates(&product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := model.Product{
		Id: uint(id),
	}

	db.DB.Delete(&product)

	return nil
}
