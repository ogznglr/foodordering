package controllers

import (
	"fmt"
	"food/helpers"
	"food/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
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
	if old_cart == "" {
		new_cart = fmt.Sprintf("%s,%s,%s", restaurant_name, product_name, piece_number)
	} else {
		cart := strings.Split(old_cart, ",")
		if cart[0] == restaurant.Slug {
			new_cart = fmt.Sprintf("%s,%s,%s", old_cart, product_name, piece_number)
		} else {
			new_cart = old_cart
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:    "cart-session",
		Value:   new_cart,
		Expires: time.Now().Add(5 * time.Minute),
	})
	return c.RedirectBack("/restaurants")
}
