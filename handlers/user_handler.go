package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gnh374/go-ld-demo/database"
	"github.com/gnh374/go-ld-demo/models"
)

// Get All Users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	// Parsing body request ke struct user
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body", "message": err.Error()})
	}

	// Simpan user ke database
	database.DB.Create(&user)
	return c.Status(201).JSON(user)
}
