package database

import (
	"fmt"
	"wsappexcel/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=wsappexcel port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to internal system database")
	}
	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{},
		&models.Product{}, &models.Order{}, &models.OrderItem{}, &models.Message{})
	fmt.Println(database)

}
