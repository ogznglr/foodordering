package middlewares

import (
	"food/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ogznglr/session"
)

var secretKey = "SecretKey"

func AuthMiddleware(c *fiber.Ctx) error {
	uid, err := IssuerMiddleware(c, secretKey)
	if err != nil {
		return err
	}

	user, err := models.User{}.First(uid)
	if err != nil {
		return err
	}
	//if the user is valid, reutrn the user and nil error
	c.Locals("user", user)

	return c.Next()
}

func IssuerMiddleware(c *fiber.Ctx, secretKey string) (int, error) {
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
