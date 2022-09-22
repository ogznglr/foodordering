package controllers

import (
	"fmt"
	"food/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ogznglr/session"
)

func NewProfilePicture(c *fiber.Ctx) error {
	s := session.New()
	issuer, err := s.Get(c, secretKey)
	if err != nil {
		return c.Redirect("/")
	}
	userid, err := strconv.Atoi(issuer)
	if err != nil {
		session.SetFlash(c, "Problem Occuerd!")
		return c.Redirect("/")
	}
	user, err := models.User{}.First(userid)
	if err != nil {
		session.SetFlash(c, "User not found!")
		return c.Redirect("/")
	}
	if user.Role != "Restaurant" {
		session.SetFlash(c, "No permission")
		return c.Redirect("/")
	}
	//the user is a valid restaurant

	picture, err := c.FormFile("picture")
	if err != nil {
		session.SetFlash(c, "File not found!")
		return c.Redirect("/myrestaurant")
	}
	err = c.SaveFile(picture, fmt.Sprintf("./uploads/%s", picture.Filename))
	if err != nil {
		session.SetFlash(c, "Problem Occured!")
		return c.Redirect("/myrestaurant")
	}
	err = user.UpdateProfilePicture(fmt.Sprintf("/uploads/%s", picture.Filename))
	if err != nil {
		session.SetFlash(c, "Couldn't update!")
		return c.Redirect("/myrestaurant")
	}

	session.SetFlash(c, "Profile Picture Saved Successfully!")
	return c.Redirect("/myrestaurant")
}
