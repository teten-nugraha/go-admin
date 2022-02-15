package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
	"strconv"
)

func AllRoles(c *fiber.Ctx) error {

	var roles []model.Role

	db.DB.Preload("Permission").Find(&roles)

	return c.JSON(&roles)
}

func CreateRole(c *fiber.Ctx) error {

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	permissionList := roleDto["permissions"].([]interface{})
	permissions := make([]model.Permission, len(permissionList))
	for i, permissionId := range permissionList {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = model.Permission{
			Id: uint(id),
		}
	}

	role := model.Role{
		Name:       roleDto["name"].(string),
		Permission: permissions,
	}

	db.DB.Create(&role)

	return c.JSON(role)
}
func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := model.Role{
		Id: uint(id),
	}

	db.DB.Preload("Permission").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	permissionList := roleDto["permissions"].([]interface{})
	permissions := make([]model.Permission, len(permissionList))
	for i, permissionId := range permissionList {
		id, _ := strconv.Atoi(permissionId.(string))

		permissions[i] = model.Permission{
			Id: uint(id),
		}
	}

	// delete existing role_permissions
	var emptyPermissions []model.Permission
	db.DB.Table("role_permission").Where("role_id", id).Delete(&emptyPermissions)

	// update role with new permission
	role := model.Role{
		Id:         uint(id),
		Name:       roleDto["name"].(string),
		Permission: permissions,
	}

	db.DB.Model(&role).Updates(role)

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
