package routes

import (
	"food/controllers"
	"food/controllers/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Listen(app *fiber.App) {
	//un-auth Gets
	app.Get("/", controllers.MainPage{}.Index)
	app.Get("/newrestaurant", controllers.NewRestaurantPage{}.Index)
	app.Get("/logout", controllers.Logout)

	//auth Gets
	app.Get("/restaurants", middlewares.AuthMiddleware, controllers.RestaurantsPage{}.Index)
	app.Get("/newaddress", middlewares.AuthMiddleware, controllers.NewAddressPage{}.Index)
	app.Get("/myrestaurant", middlewares.AuthMiddleware, controllers.MyRestaurantPage{}.Index)
	app.Get("/newproduct", middlewares.AuthMiddleware, controllers.NewProductPage{}.Index)
	app.Get("/product/edit/:data", middlewares.AuthMiddleware, controllers.EditProductPage{}.Index)
	app.Get("/restaurant/:data", middlewares.AuthMiddleware, controllers.RestaurantPage{}.Index)
	app.Get("/restaurant/:restaurant/:product", middlewares.AuthMiddleware, controllers.ProductDetailPage{}.Index)
	app.Get("/product/delete/:data", middlewares.AuthMiddleware, controllers.DeleteProduct)
	app.Get("/cart", middlewares.AuthMiddleware, controllers.CartPage{}.Index)
	app.Get("/cart/delete/:data", middlewares.AuthMiddleware, controllers.DeleteFromCart)

	//un-auth Posts
	app.Post("/login", controllers.Login)
	app.Post("/newuser", controllers.NewUser)
	app.Post("/newrestaurant", controllers.NewRestaurant)

	//auth Posts
	app.Post("/newaddress", middlewares.AuthMiddleware, controllers.NewAddress)
	app.Post("/newproduct", middlewares.AuthMiddleware, controllers.NewProduct)
	app.Post("/newprofilepicture", middlewares.AuthMiddleware, controllers.NewProfilePicture)
	app.Post("/product/edit", middlewares.AuthMiddleware, controllers.EditProduct)
	app.Post("/addtocart", middlewares.AuthMiddleware, controllers.Addtocart)

	app.Static("/assets/", "./view/assets/")
	app.Static("/uploads/", "./uploads/")
}
