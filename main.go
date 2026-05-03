package main

import (
	"log"

	"go-fiber-crud/config"
	"go-fiber-crud/database"
	"go-fiber-crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadEnv()
	db := database.ConnectDB()

	app := fiber.New()
	app.Use(logger.New())

	routes.RegisterTaskRoutes(app, db)

	port := config.GetEnv("APP_PORT", "3000")
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}