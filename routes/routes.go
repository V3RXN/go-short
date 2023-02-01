package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/v3rxn/go-short/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/link", controllers.CreateURL)
	app.Get("/:id", controllers.Redirect)
	// app.Get("/:id", func (c *fiber.Ctx) error {
	// 	return c.SendString(c.Params("id"))
	// })
}