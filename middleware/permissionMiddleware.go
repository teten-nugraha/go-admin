package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/teten-nugraha/go-admin/db"
	"github.com/teten-nugraha/go-admin/model"
	"github.com/teten-nugraha/go-admin/util"
	"strconv"
)

func IsAuthorized(c *fiber.Ctx, page string) error {
	cookie := c.Cookies("jwt")

	id, err := util.ParseJwt(cookie)
	if err != nil {
		return err
	}

	userId, _ := strconv.Atoi(id)

	user := model.User{
		Id: uint(userId),
	}

	db.DB.Preload("Role").Find(&user)

	role := model.Role{
		Id: user.RoleId,
	}

	if c.Method() == "GET" {
		for _, permission := range role.Permission {
			if permission.Name == "view_"+page || permission.Name == "edit_"+page {
				return nil
			}
		}
	} else {
		for _, permission := range role.Permission {
			if permission.Name == "edit"+page {
				return nil
			}
		}
	}

	c.Status(fiber.StatusUnauthorized)
	return errors.New("unauthorized")
}
