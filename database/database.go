package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	str := "root:password@tcp(db:3306)/foodorder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database")
	}
	DB = db
}

func Migrate(d interface{}) error {

	err := DB.AutoMigrate(d)
	if err != nil {
		panic("Database couldn't be updated")
	}
	return nil
}
