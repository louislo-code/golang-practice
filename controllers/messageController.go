package controllers

import (
	"strconv"
	"strings"
	"wsappexcel/database"
	"wsappexcel/models"

	// "github.com/xuri/excelize/v2"

	"github.com/gofiber/fiber/v2"
)

func AllMessages(c *fiber.Ctx) error {
	// if err := middlewares.IsAuthorized(c, "messages"); err != nil {
	// 	return err
	// }

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Message{}, page))
}

type UnformatedText struct {
	text string
}

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func CreateMessage(c *fiber.Ctx) error {
	// if err := middlewares.IsAuthorized(c, "messages"); err != nil {
	// 	return err
	// }
	// if err := middlewares.IsAuthorized(c, "messages"); err != nil {
	// 	return err
	// }

	var messages models.Message

	if err := c.BodyParser(&messages); err != nil {
		return err
	}

	database.DB.Create(&messages)

	return c.JSON(messages)

}

func GetMessage(c *fiber.Ctx) error {
	// if err := middlewares.IsAuthorized(c, "users"); err != nil {
	// 	return err
	// }

	id, _ := strconv.Atoi(c.Params("id"))

	message := models.Message{
		Id: uint(id),
	}
	database.DB.Find(&message)

	return c.JSON(message)
}

func UpdateMessage(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))
	message := models.Message{
		Id: uint(id),
	}

	if err := c.BodyParser(&message); err != nil {
		return err
	}

	database.DB.Model(&message).Updates(message)

	return c.JSON(message)
}

func DeleteMessage(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	message := models.Message{
		Id: uint(id),
	}

	database.DB.Delete(&message)

	return nil
}
