package handlers

import (
	"github.com/gnh374/go-ld-demo/config"
	"github.com/gnh374/go-ld-demo/database"
	"github.com/gnh374/go-ld-demo/models"
	"github.com/gofiber/fiber/v2"
	"github.com/launchdarkly/go-sdk-common/v3/ldcontext"
)

// Get All Users
func GetUsers(c *fiber.Ctx) error {// Transfer API: Mengurangi saldo user

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

func Transfer(c *fiber.Ctx) error {


	var LDClient = config.LDClient

	// Ambil ID user dari parameter URL
	id := c.Params("id_user")


	// Ambil user dari database berdasarkan ID
	var user models.User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	userContext := ldcontext.NewBuilder(id).
	SetString("name", user.Name).
	SetString("email", user.Email).
	Build()

	flagKey := "payment-flag"
	showFeature, err := LDClient.BoolVariation(flagKey, userContext, false)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to evaluate feature flag", "message": err.Error()})
	}


	// Parsing request body untuk mendapatkan jumlah transfer
	var request struct {
		Amount float64 `json:"amount"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body", "message": err.Error()})
	}

	if(showFeature){
		// Validasi jika saldo cukup
		if user.Savings < request.Amount {
			return c.Status(400).JSON(fiber.Map{"error": "Insufficient balance"})
		}

		// Kurangi saldo user
		user.Savings -= request.Amount
		database.DB.Save(&user)

		// Return response
		return c.Status(200).JSON(fiber.Map{
			"message": "Transfer successful",
			"user":    user,
		})
	}

	return c.Status(500).JSON(fiber.Map{"message" :"Feature off, try again later"})

	
}

