package controllers

import (
	"strconv"
	"wsappexcel/database"
	"wsappexcel/models"

	"github.com/gofiber/fiber/v2"
)

func AllRoles(c *fiber.Ctx) error {
	var role []models.Role

	database.DB.Find(&role)

	return c.JSON(role)
}

// type RoleCreateDTO struct {
// 	name        string
// 	permissions []string
// }

func CreateRole(c *fiber.Ctx) error {

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	// var permission123 string = database.DB.Select("Name").Where("name = 'edit_users'").Find(&permissions)
	// var permission123id string = database.DB.Select("Id").Where("name = 'edit_users'").Find(&permissions)

	// superiorUser := models.Role{
	// 	Name:        "superior_user",
	// 	Permissions: [{permipermission123id,permission123}],
	// }
	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}
	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDto fiber.Map

	if err := c.BodyParser(&roleDto); err != nil {
		return err
	}

	list := roleDto["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		id, _ := strconv.Atoi(permissionId.(string))
		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	var result []interface{}
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto["name"].(string),
		Permissions: permissions,
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return nil
}
