package models

import (
	"fmt"
	"food/database"
)

type Address struct {
	ID            int
	UserID        int
	User          User
	District      string
	City          string
	Neighbourhood string
	Street        string
	BuildingNo    string
	DoorNo        string
	Lat           float64
	Lng           float64
}

func (address Address) Delete() {
	database.DB.Unscoped().Delete(&address)
}

func (address Address) New() {
	database.DB.Create(&address)
}

func (address Address) First(userid int) (Address, error) {
	db := database.DB.Where("user_id = ?", userid).First(&address)
	if db.Error != nil {
		fmt.Println("adres bulunamadı")
		return Address{}, db.Error
	}
	return address, nil
}
