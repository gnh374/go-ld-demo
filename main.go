package main

import (
	"github.com/gnh374/go-ld-demo/config"
	"github.com/gnh374/go-ld-demo/database"
	"github.com/gnh374/go-ld-demo/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Enable CORS for all origins
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://13.212.48.92:3000",  // Mengizinkan semua origin
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",  // Mengizinkan metode tertentu
		AllowHeaders: "Content-Type, Authorization",  // Header yang diizinkan
		AllowCredentials: true,  // Mengizinkan kredensial
	}))
	

	// Connect Database
	database.ConnectDB()
	database.MigrateDB()

	config.InitiateLDClient()
	defer config.CloseLDClient()

	// Routes
	app.Get("/users", handlers.GetUsers)
	app.Post("/users", handlers.CreateUser)
	app.Post("/transfer/:id_user", handlers.Transfer)

	// Start server
	app.Listen(":3000")
}
