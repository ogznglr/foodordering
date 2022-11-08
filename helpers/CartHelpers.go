package helpers

import (
	"fmt"
	"food/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const product_property_number = 3

func CartToModels(c *fiber.Ctx) ([]models.Response_cart, string, error) {
	cart_session := c.Cookies("cart-session")
	cart := strings.Split(cart_session, ",")
	restaurant_name := cart[0]
	var response_cart []models.Response_cart
	//check if restaurant from cookie, is valid.
	restaurant, err := models.User{}.FirstRestaurantWithSlug(restaurant_name)
	if err != nil {
		return nil, "", err
	}
	//take all the product properties from cookie and add to response cart as a real product.
	for i := 1; i < len(cart); i += product_property_number {
		uuid := cart[i]
		product, err := models.Product{}.FirstWithSlug(cart[i+1])
		entity, _ := strconv.Atoi(cart[i+2])

		//did we find the product and has the product the same restaurantid with the restaurant
		if err != nil || product.RestaurantID != strconv.Itoa(int(restaurant.ID)) {
			continue
		}
		//is there any problem with converting the price from string to float
		unit_price, err := strconv.ParseFloat(product.Price, 64)
		if err != nil {
			continue
		}

		response_cart = append(response_cart, models.Response_cart{
			ID:      uuid,
			Product: product,
			Entity:  entity,
			Price:   fmt.Sprintf("%.2f", unit_price*float64(entity)),
		})
	}

	return response_cart, restaurant_name, nil
}

func ModelsToCart(restaurant_slug string, models []models.Response_cart) string {
	cart := fmt.Sprintf("%s", restaurant_slug)
	for _, value := range models {
		cart = fmt.Sprintf("%s,%s,%s,%d", cart, value.ID, value.Product.Slug, value.Entity)
	}
	fmt.Println("models returned to cart: ", cart)
	return cart
}
