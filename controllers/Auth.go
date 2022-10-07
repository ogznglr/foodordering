package controllers

import (
	"fmt"
	"food/database"
	"food/helpers"
	"food/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/ogznglr/session"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "SecretKey"

const cost = 10

func NewUser(c *fiber.Ctx) error {
	fn := c.FormValue("firstname")
	ln := c.FormValue("lastname")
	email := c.FormValue("email")
	pw := c.FormValue("password")

	//is email valid?
	email, err := helpers.ValidEmail(email)
	if err != nil {
		session.SetFlash(c, "Please enter a valid email!")
		return c.Redirect("/")
	}

	//is there a user with this email?
	var user models.User
	db := database.DB.Where("email = ?", email).First(&user)
	if db.Error == nil {
		session.SetFlash(c, "There is already a user with this email. Please Login!")
		return c.Redirect("/")
	}

	//Hash the password before saving into the database.
	password, err := bcrypt.GenerateFromPassword([]byte(pw), cost)
	if err != nil {
		return err
	}
	pw = string(password)

	db = database.DB.Create(&models.User{
		FirstName: fn,
		LastName:  ln,
		Email:     email,
		Password:  pw,
		Role:      "User",
	})
	if db.Error != nil {
		fmt.Println(err)
		return db.Error
	}
	return c.Redirect("/")
}

func NewRestaurant(c *fiber.Ctx) error {
	nm := c.FormValue("firstname")
	fmt.Println("first name : ", nm)
	email := c.FormValue("email")
	pw := c.FormValue("password")

	//is email valid?
	email, err := helpers.ValidEmail(email)
	if err != nil {
		session.SetFlash(c, "Please enter a valid email!")
		return c.Redirect("/")
	}

	//is there a user with this email?
	var user models.User
	db := database.DB.Where("email = ?", email).First(&user)
	if db.Error == nil {
		session.SetFlash(c, "There is already a user with this email. Please Login!")
		return c.Redirect("/")
	}

	//Hash the password before saving into the database.
	password, err := bcrypt.GenerateFromPassword([]byte(pw), cost)
	if err != nil {
		return err
	}
	pw = string(password)

	db = database.DB.Create(&models.User{
		FirstName: nm,
		Email:     email,
		Password:  pw,
		Role:      "Restaurant",
		Slug:      slug.Make(nm),
	})
	if db.Error != nil {
		fmt.Println(err)
		return db.Error
	}
	return c.Redirect("/")
}

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	var user models.User
	database.DB.Where("email = ?", email).First(&user)
	//is email true?
	if user.ID == 0 {
		session.SetFlash(c, "Username or Password is wrong!")
		return c.Redirect("/")
	}
	//is password true?
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		session.SetFlash(c, "Username or Password is wrong!")
		return c.Redirect("/")
	}
	s := session.New(12)
	s.Set(c, strconv.Itoa(int(user.ID)), secretKey)

	session.SetFlash(c, "Login successfully!")
	return c.Redirect("/")
}
func Logout(c *fiber.Ctx) error {
	s := session.New()
	s.Delete(c)
	session.SetFlash(c, "Loged out Successfully")
	return c.Redirect("/")
}
func NewAddress(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	currentaddr, err := models.Address{}.First(int(user.ID))
	//if there is an address, delete it first
	if err == nil {
		currentaddr.Delete()
	}

	address := &models.Address{
		UserID:        int(user.ID),
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
	user.UpdateCity(address.City)

	session.SetFlash(c, "Address has been added successfully!")
	return c.Redirect("/newaddress")
}
