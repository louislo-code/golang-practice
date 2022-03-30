package models

import "gorm.io/gorm"

type Message struct {
	Id       uint   `json:"id"`
	Date     string `json:"date"`
	Name     string `json:"name"`
	Product  string `json:"product"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
}

func (message *Message) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Message{}).Count(&total)

	return total
}

func (message *Message) Take(db *gorm.DB, limit int, offset int) interface{} {
	var messages []Message

	db.Offset(offset).Limit(limit).Find(&messages)

	return messages
}
