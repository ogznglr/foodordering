package database

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var RClient *redis.Client

func Connection() {
	str := "root:@tcp(localhost:3306)/foodorder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database")
	}
	DB = db

	//Redis Connection
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       10,
	})

	if client == nil {
		panic("Couldn't connect to the Ratabase")
	}
	RClient = client
}

func Migrate(d interface{}) error {

	err := DB.AutoMigrate(d)
	if err != nil {
		panic("Database couldn't be updated")
	}
	return nil
}
