package controllers

import (
	"fmt"
	"food/helpers"
	"food/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ogznglr/session"
)

func Addtocart(c *fiber.Ctx) error {

	restaurant_name := c.FormValue("restaurant-name")
	user := c.Locals("user").(models.User)
	user_address, _ := models.Address{}.First(int(user.ID))
	restaurant, err := models.User{}.FirstRestaurantWithSlug(restaurant_name)
	restaurant_address, _ := models.Address{}.First(int(restaurant.ID))

	old_cart := c.Cookies("cart-session")
	product_name := c.FormValue("product-name")
	piece_number := c.FormValue("piece-number")
	var new_cart string

	//if the user's role is okay for adding cart
	if user.Role != "User" {
		session.SetFlash(c, "Unauthorized")
		return c.Redirect("/")
	}

	//is restaurant distance okey for ordering this product?
	distance, err := helpers.GetDistance(user_address, restaurant_address)
	if err != nil || distance > 10000 {
		session.SetFlash(c, "No such restaurant")
		return c.Redirect(fmt.Sprintf("/restaurant/%s", restaurant_name))
	}

	//Does restaurant have this product really?
	product, err := models.Product{}.FirstWithSlug(product_name)
	if err != nil || product.RestaurantID != fmt.Sprintf("%d", restaurant.ID) {
		session.SetFlash(c, "No such product")
		return c.Redirect(fmt.Sprintf("/restaurant/%s", restaurant_name))
	}

	//Restaurant has the product, so we can add to cart safely!

	//this uuid will be the id of product in the cart
	newuuid := uuid.NewString()

	if old_cart == "" {
		new_cart = fmt.Sprintf("%s,%s,%s,%s", restaurant_name, newuuid, product_name, piece_number)
	} else {
		cart := strings.Split(old_cart, ",")
		if cart[0] == restaurant.Slug {
			new_cart = fmt.Sprintf("%s,%s,%s,%s", old_cart, newuuid, product_name, piece_number)
		} else {
			//if restaurant in the cart is not same with the restaurant we want to order, dont change anything.
			new_cart = old_cart
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:    "cart-session",
		Value:   new_cart,
		Expires: time.Now().Add(2 * time.Hour),
	})
	return c.Redirect(fmt.Sprintf("/restaurant/%s", restaurant.Slug))
}

func DeleteFromCart(c *fiber.Ctx) error {
	ID := c.Params("data")
	parsed_cart, restaurant_name, err := helpers.CartToModels(c)
	if err != nil {
		session.SetFlash(c, "Cart Error Happened")
		return c.Redirect("/cart")
	}

	var new_cart []models.Response_cart
	for _, value := range parsed_cart {
		if value.ID != ID {
			new_cart = append(new_cart, value)
		}
	}
	new_cart_string := helpers.ModelsToCart(restaurant_name, new_cart)
	if new_cart == nil {
		c.Cookie(&fiber.Cookie{
			Name:    "cart-session",
			Expires: time.Now().Add(-10 * time.Hour),
		})
	} else {
		c.Cookie(&fiber.Cookie{
			Name:    "cart-session",
			Value:   new_cart_string,
			Expires: time.Now().Add(2 * time.Hour),
		})
	}

	return c.Redirect("/cart")
}
