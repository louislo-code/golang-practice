package controllers

import (
	"wsappexcel/database"
	"wsappexcel/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission
	database.DB.Find(&permissions)
	// permission111 := models.Permission{Id: 15, Name: "trying"}
	// database.DB.Create(&permission111)
	return c.JSON(permissions)

}
