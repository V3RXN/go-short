package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/v3rxn/go-short/config"
	"github.com/v3rxn/go-short/database"
	"github.com/v3rxn/go-short/models"
	"github.com/v3rxn/go-short/utils"
)

func CreateURL(c *fiber.Ctx) error {
	
	var data map[string]string
	var link models.Link

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	link = models.Link{
		Id: 	utils.GenerateUniqueId(),	
		URL: 	data["url"],
	}

	database.DB.Create(&link)
	
	shortURL := fmt.Sprintf("http://%v/%v", config.Domain, link.Id)

	return c.Status(fiber.StatusCreated).JSON(shortURL)
}

func Redirect(c *fiber.Ctx) error {
	id := c.Params("id")

	var link models.Link
	if err := database.DB.Where("id = ?", id).First(&link).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Could not find the link",
		})
	}

	return c.Redirect(link.URL, fiber.StatusPermanentRedirect)
}
