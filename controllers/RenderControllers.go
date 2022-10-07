package controllers

import (
	"food/helpers"
	"food/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ogznglr/session"
)

type MainPage struct {
}
type NewRestaurantPage struct {
}
type NewAddressPage struct {
}
type RestaurantsPage struct {
}
type MyRestaurantPage struct {
}
type NewProductPage struct {
}
type EditProductPage struct {
}
type RestaurantPage struct {
}
type ProductDetailPage struct {
}

func (MainPage) Index(c *fiber.Ctx) error {

	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		alert := session.GetFlash(c)
		return c.Render("mainpage", fiber.Map{
			"alert": alert,
		})
	}

	if user.Role == "User" {
		return c.Redirect("/restaurants")
	}
	if user.Role == "Restaurant" {
		return c.Redirect("/myrestaurant")
	}
	alert := session.GetFlash(c)
	return c.Render("mainpage", fiber.Map{
		"alert": alert,
	})
}
func (NewRestaurantPage) Index(c *fiber.Ctx) error {
	alert := session.GetFlash(c)
	return c.Render("newrestaurant", fiber.Map{
		"alert": alert,
	})
}
func (RestaurantsPage) Index(c *fiber.Ctx) error {
	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "Please Login!")
		return c.Redirect("/")
	}

	if user.Role != "User" {
		session.SetFlash(c, "You have no permission!")
		return c.Redirect("/")
	}
	//Does user have a valid address
	useraddress, err := models.Address{}.First(int(user.ID))
	if err != nil {
		session.SetFlash(c, "Please enter a valid address")
		return c.Redirect("/newaddress")
	}
	//Get the restaurants from your city.
	restaurants := models.User{}.FindRestaurants(useraddress.City)
	if restaurants == nil {
		session.SetFlash(c, "There is no any restaurant close to you!")
		return c.Redirect("/newaddress")
	}

	restaurants_slice := []models.Response_restaurants{}

	for _, element := range restaurants {
		restaddress, _ := models.Address{}.First(int(element.ID))
		dist, err := helpers.GetDistance(useraddress, restaddress)
		dist = dist / 1000
		if err != nil {
			continue
		}

		restaurants_slice = append(restaurants_slice, models.Response_restaurants{
			Restaurant: element,
			Address:    restaddress,
			Distance:   dist,
		})
	}

	//This function sorts the restaurants from lowest distance to highest.
	restaurants_slice = helpers.QuickSort(restaurants_slice, 0, len(restaurants_slice))

	alert := session.GetFlash(c)
	return c.Render("usermainpage", fiber.Map{
		"restaurants": restaurants_slice,
		"user":        user,
		"alert":       alert,
	})
}
func (NewAddressPage) Index(c *fiber.Ctx) error {
	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "Please Login!")
		return c.Redirect("/")
	}

	address, _ := models.Address{}.First(int(user.ID))

	alert := session.GetFlash(c)
	if user.Role == "Restaurant" {
		return c.Render("restaurantnewaddresspage", fiber.Map{
			"address": address,
			"user":    user,
			"alert":   alert,
		})
	} else {
		return c.Render("newaddresspage", fiber.Map{
			"address": address,
			"user":    user,
			"alert":   alert,
		})
	}
}
func (MyRestaurantPage) Index(c *fiber.Ctx) error {

	restaurant, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "No such user!")
		return c.Redirect("/")
	}

	if restaurant.Role != "Restaurant" {
		session.SetFlash(c, "No Permission!")
		return c.Redirect("/")
	}
	//if restaurant has no address, redirect to newaddress page.
	addr, _ := models.Address{}.First(int(restaurant.ID))
	if addr.ID == 0 {
		session.SetFlash(c, "You have no address!")
		return c.Redirect("/newaddress")
	}

	//user is a valid restaurant, there is no problem.
	//get all the products the restaurant has
	products, _ := models.Product{}.Find(int(restaurant.ID))

	alert := session.GetFlash(c)
	return c.Render("restaurantmainpage", fiber.Map{
		"user":     restaurant,
		"products": products,
		"address":  addr,
		"alert":    alert,
	})
}
func (NewProductPage) Index(c *fiber.Ctx) error {

	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "User not found!")
		return c.Redirect("/")
	}

	if user.Role != "Restaurant" {
		session.SetFlash(c, "No permission")
		return c.Redirect("/")
	}
	//user is a valid restaurtan, there is no problem
	alert := session.GetFlash(c)
	return c.Render("newproductpage", fiber.Map{
		"alert": alert,
	})
}
func (EditProductPage) Index(c *fiber.Ctx) error {

	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "User not found!")
		return c.Redirect("/")
	}

	if user.Role != "Restaurant" {
		session.SetFlash(c, "No permission")
		return c.Redirect("/")
	}
	//user is a valid restaurtan, there is no problem

	product, err := helpers.ProductValidation(c.Params("data"))
	if err != nil {
		session.SetFlash(c, "Product has not found!")
		return c.Redirect("/myrestaurant")
	}

	//is the product owned by this restaurant?
	if strconv.Itoa(int(user.ID)) != product.RestaurantID {
		session.SetFlash(c, "Permission Error!")
		return c.Redirect("/myrestaurant")
	}

	//if nothing wrong, render edit product page.
	alert := session.GetFlash(c)
	return c.Render("editproductpage", fiber.Map{
		"alert":   alert,
		"product": product,
		"user":    user,
	})
}
func (RestaurantPage) Index(c *fiber.Ctx) error {

	//User Validation process
	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "Please Login!")
		return c.Redirect("/")
	}

	if user.Role != "User" {
		session.SetFlash(c, "No permission!")
		return c.Redirect("/")
	}

	//user is a valid User, there is no problem.

	restaurantname := c.Params("data")
	restaurant, err := models.User{}.FirstRestaurantWithSlug(restaurantname)

	if err != nil {
		session.SetFlash(c, "Restaurant's not found!")
		return c.Redirect("/restaurants")
	}
	address, _ := models.Address{}.First(int(restaurant.ID))
	products, _ := models.Product{}.Find(int(restaurant.ID))

	alert := session.GetFlash(c)
	return c.Render("restaurantcontentpage", fiber.Map{
		"alert":      alert,
		"restaurant": restaurant,
		"user":       user,
		"address":    address,
		"products":   products,
	})
}
func (ProductDetailPage) Index(c *fiber.Ctx) error {
	//User Validation process
	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "Please Login!")
		return c.Redirect("/")
	}

	if user.Role != "User" {
		session.SetFlash(c, "No permission!")
		return c.Redirect("/")
	}
	//user is a valid User, there is no problem.

	rslug := c.Params("restaurant")
	pslug := c.Params("product")

	restaurant, _ := models.User{}.FirstRestaurantWithSlug(rslug)
	product, _ := models.Product{}.FirstWithSlug(pslug)

	userAddress, _ := models.Address{}.First(int(user.ID))
	restaurantAddress, _ := models.Address{}.First(int(restaurant.ID))

	distance, _ := helpers.GetDistance(restaurantAddress, userAddress)

	if distance >= 10000 {
		session.SetFlash(c, "The distance is too much. Please select another restaurant!")
		return c.Redirect("/restaurants")
	}

	alert := session.GetFlash(c)
	return c.Render("productdetailpage", fiber.Map{
		"restaurant": restaurant,
		"product":    product,
		"alert":      alert,
		"user":       user,
	})
}
