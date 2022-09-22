package models

import (
	"food/database"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string
	Description  string
	PictureURL   string
	RestaurantID string
	Restaurant   User
	Price        string
	Slug         string
}

func (product Product) Find(restaurantid int) ([]Product, error) {
	var products []Product
	db := database.DB.Where("restaurant_id = ?", restaurantid).Find(&products)
	if db.Error != nil {
		return products, db.Error
	}
	return products, nil
}

func (product Product) First(productid int) Product {
	database.DB.First(&product, productid)
	return product
}

func (product Product) New() error {
	db := database.DB.Create(&product)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (product Product) Delete() {
	database.DB.Unscoped().Delete(&product)
}

func (product Product) Update(newproduct *Product) error {
	db := database.DB.Model(&product).Updates(newproduct)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (product Product) FirstWithSlug(slug string) (Product, error) {
	db := database.DB.Where("Slug = ?", slug).First(&product)
	if db.Error != nil {
		return product, db.Error
	}
	return product, nil
}
