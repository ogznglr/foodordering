package helpers

import (
	"food/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ogznglr/session"
)

func UserValidation(c *fiber.Ctx, secretKey string) (models.User, error) {
	s := session.New()
	issuer, err := s.Get(c, secretKey)
	if err != nil {
		return models.User{}, err
	}

	uid, err := strconv.Atoi(issuer)
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
