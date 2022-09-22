package models

import (
	"food/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role       string
	FirstName  string
	LastName   string
	Password   string
	Email      string `gorm:"unique"`
	City       string
	PictureURL string
	Slug       string
}

func (user User) New() error {
	db := database.DB.Create(&user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
func (user User) First(userid int) (User, error) {
	db := database.DB.First(&user, userid)
	if db.Error != nil {
		return User{}, db.Error
	}
	return user, nil
}
func (user User) FindRestaurants(city string) []User {
	var users []User
	database.DB.Where("city = ? AND role = ?", city, "Restaurant").Find(&users)
	if len(users) == 0 {
		return nil
	}
	return users
}
func (user User) FirstRestaurantWithSlug(slug string) (User, error) {

	db := database.DB.Where("slug = ?", slug).First(&user)
	if db.Error != nil {
		return user, db.Error
	}
	return user, nil
}
func (user User) UpdateCity(city string) error {
	db := database.DB.Model(&user).Update("city", city)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
func (user User) UpdateProfilePicture(url string) error {
	db := database.DB.Model(&user).Update("picture_url", url)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
