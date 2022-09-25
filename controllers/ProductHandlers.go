package controllers

import (
	"fmt"
	"food/helpers"
	"food/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/ogznglr/session"
)

func NewProduct(c *fiber.Ctx) error {
	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "User not found!")
		return c.Redirect("/")
	}

	if user.Role != "Restaurant" {
		session.SetFlash(c, "No permission")
		return c.Redirect("/")
	}
	//user is a valid restaurant

	name := c.FormValue("name")
	description := c.FormValue("description")
	price := c.FormValue("price")
	picture, err := c.FormFile("picture")

	if err != nil {
		session.SetFlash(c, "Problem Occured!")
		return c.Redirect("/")
	}
	if picture == nil {
		session.SetFlash(c, "You must add a picture!")
		return c.Redirect("/newproduct")
	}

	err = c.SaveFile(picture, fmt.Sprintf("./uploads/%s", picture.Filename))

	if err != nil {
		return err
	}

	product := models.Product{
		Restaurant:  user,
		Name:        name,
		Description: description,
		Price:       price,
		PictureURL:  fmt.Sprintf("/uploads/%s", picture.Filename),
		Slug:        slug.Make(name),
	}
	product.New() //add to database

	session.SetFlash(c, "The product has been added successfully!")
	return c.Redirect("/myrestaurant")
}

func DeleteProduct(c *fiber.Ctx) error {
	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "User not found!")
		return c.Redirect("/")
	}

	if user.Role != "Restaurant" {
		session.SetFlash(c, "No permission")
		return c.Redirect("/")
	}
	//user is a valid restaurant

	id := c.Params("data")

	productid, err := strconv.Atoi(id)
	if err != nil {
		session.SetFlash(c, "Problem Occured!")
		return c.Redirect("/myrestaurant")
	}
	var product models.Product
	product = product.First(productid)
	//is the product owned by this restaurant?

	if strconv.Itoa(int(user.ID)) != product.RestaurantID {
		session.SetFlash(c, "Problem Occured!")
		return c.Redirect("/myrestaurant")
	}

	product.Delete()

	session.SetFlash(c, "Product Deleted Successfully!")
	return c.Redirect("/myrestaurant")
}

func EditProduct(c *fiber.Ctx) error {
	user, err := helpers.UserValidation(c, secretKey)
	if err != nil {
		session.SetFlash(c, "User not found!")
		return c.Redirect("/")
	}

	if user.Role != "Restaurant" {
		session.SetFlash(c, "No permission")
		return c.Redirect("/")
	}
	//user is a valid restaurant

	id := c.FormValue("productid")

	productid, err := strconv.Atoi(id)
	if err != nil {
		session.SetFlash(c, "Problem Occured!")
		return c.Redirect("/myrestaurant")
	}
	var product models.Product
	product = product.First(productid)
	//is the product owned by this restaurant?

	if strconv.Itoa(int(user.ID)) != product.RestaurantID {
		session.SetFlash(c, "Problem Occured!")
		return c.Redirect("/myrestaurant")
	}

	var newproduct models.Product

	name := c.FormValue("name")
	description := c.FormValue("description")
	price := c.FormValue("price")
	picture, _ := c.FormFile("picture")

	//Picture setting opearations
	if picture != nil {
		err = c.SaveFile(picture, fmt.Sprintf("./uploads/%s", picture.Filename))
		if err == nil {
			newproduct.PictureURL = fmt.Sprintf("/uploads/%s", picture.Filename)
		} else {
			session.SetFlash(c, "The product hasn't been updated!")
			return c.Redirect("/myrestaurant")
		}
	} else {
		newproduct.PictureURL = product.PictureURL
	}

	newproduct.Name = name
	newproduct.Description = description
	newproduct.Price = price
	newproduct.Slug = slug.Make(name)

	err = product.Update(&newproduct)

	if err != nil {
		session.SetFlash(c, "Problem Occured!")
		return c.Redirect("/myrestaurant")
	}
	session.SetFlash(c, "The post edited successfully!")
	return c.Redirect("/myrestaurant")
}
