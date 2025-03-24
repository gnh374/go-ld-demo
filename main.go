package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gnh374/go-ld-demo/database"
	"github.com/gnh374/go-ld-demo/handlers"
)

func main() {
	app := fiber.New()

	// Connect Database
	database.ConnectDB()
	database.MigrateDB()

	// Routes
	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.CreateUser)

	// Start server
	app.Listen(":3000")
}
