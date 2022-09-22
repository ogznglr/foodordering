package controllers

import (
	"food/helpers"
	"food/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ogznglr/session"
)

func NewAddress(c *fiber.Ctx) error {
	s := session.New()
	issuer, err := s.Get(c, secretKey)
	if err != nil {
		session.SetFlash(c, "A problem occured!")
		c.Redirect("/newaddress")
	}

	userid, err := strconv.Atoi(issuer)
	if err != nil {
		session.SetFlash(c, "A problem occured!")
		c.Redirect("/newaddress")
	}

	currentaddr, err := models.Address{}.First(userid)
	//if there is an address, delete it first
	if err == nil {
		currentaddr.Delete()
	}

	address := &models.Address{
		UserID:        userid,
		District:      c.FormValue("district"),
		City:          c.FormValue("city"),
		Neighbourhood: c.FormValue("neighbourhood"),
		BuildingNo:    c.FormValue("buildingno"),
		Street:        c.FormValue("street"),
		DoorNo:        c.FormValue("doorno"),
	}
	lat, lng, err := helpers.GetCoordinates(address)

	if err != nil {
		return err
	}

	address.Lat = lat
	address.Lng = lng

	address.New()

	//Updates user's city
	user, _ := models.User{}.First(userid)
	user.UpdateCity(address.City)

	session.SetFlash(c, "Address has been added successfully!")
	return c.Redirect("/newaddress")
}
