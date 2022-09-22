package routes

import (
	"food/controllers"

	"github.com/gofiber/fiber/v2"
)

func Listen(app *fiber.App) {
	app.Get("/", controllers.MainPage{}.Index)
	app.Get("/newrestaurant", controllers.NewRestaurantPage{}.Index)
	app.Get("/restaurants", controllers.RestaurantsPage{}.Index)
	app.Get("/newaddress", controllers.NewAddressPage{}.Index)
	app.Get("/myrestaurant", controllers.MyRestaurantPage{}.Index)
	app.Get("/newproduct", controllers.NewProductPage{}.Index)
	app.Get("/logout", controllers.Logout)
	app.Get("/product/edit/:data", controllers.EditProductPage{}.Index)
	app.Get("/restaurant/:data", controllers.RestaurantPage{}.Index)
	app.Get("/restaurant/:restaurant/:product", controllers.ProductDetailPage{}.Index)

	app.Get("/product/delete/:data", controllers.DeleteProduct)

	//post requests
	app.Post("/newuser", controllers.NewUser)
	app.Post("/newrestaurant", controllers.NewRestaurant)
	app.Post("/login", controllers.Login)
	app.Post("/newaddress", controllers.NewAddress)
	app.Post("/newproduct", controllers.NewProduct)
	app.Post("/newprofilepicture", controllers.NewProfilePicture)
	app.Post("/product/edit", controllers.EditProduct)

	app.Static("/assets/", "./view/assets/")
	app.Static("/uploads/", "./uploads/")
}
