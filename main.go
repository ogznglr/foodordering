package main

import (
	"food/database"
	"food/models"
	"food/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() {
	database.Connection()

	database.Migrate(models.User{})
	database.Migrate(models.Order{})
	database.Migrate(models.Product{})
	database.Migrate(models.Address{})
}

func main() {
	engine := html.New("./view", ".html")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Listen(app)

	app.Listen(":80")

}
