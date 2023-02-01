package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/v3rxn/go-short/config"
	"github.com/v3rxn/go-short/database"
	"github.com/v3rxn/go-short/routes"
)

func main() {

	config.LoadEnv()
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(config.Port))
	
}