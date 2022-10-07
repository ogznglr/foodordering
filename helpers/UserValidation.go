package helpers

import (
	"food/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ogznglr/session"
)

func UserValidation(c *fiber.Ctx, secretKey string) (models.User, error) {
	uid, err := IssuerToId(c, secretKey)
	if err != nil {
		return models.User{}, err
	}

	user, err := models.User{}.First(uid)
	if err != nil {
		return models.User{}, err
	}
	//if the user is valid, reutrn the user and nil error
	return user, nil
}

func IssuerToId(c *fiber.Ctx, secretKey string) (int, error) {
	s := session.New()
	issuer, err := s.Get(c, secretKey)
	if err != nil {
		return 0, err
	}

	uid, err := strconv.Atoi(issuer)
	if err != nil {
		return 0, err
	}
	//user id from access token is converted to integer successfully.
	return uid, nil
}
