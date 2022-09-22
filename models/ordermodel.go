package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID       string
	User         User
	RestaurantID string
	Restaurant   User
	ProductID    string
	Product      Product
	Quantity     byte
}
