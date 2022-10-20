package helpers

import (
	"food/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetCart(c *fiber.Ctx) ([]models.Response_cart, error) {
	cart_session := c.Cookies("cart-session")
	cart := strings.Split(cart_session, ",")
	restaurant_name := cart[0]
	var response_cart []models.Response_cart
	//check if restaurant from cookie, is valid.
	_, err := models.User{}.FirstRestaurantWithSlug(restaurant_name)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(cart); i += 2 {
		entity, _ := strconv.Atoi(cart[i+1])
		response_cart = append(response_cart, models.Response_cart{
			ProductName: cart[i],
			Entity:      entity,
		})
	}
	return response_cart, nil

}
