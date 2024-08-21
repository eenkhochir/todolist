package handlers

import (
	"todolist-backend/database"
	"todolist-backend/models"

	"github.com/gofiber/fiber/v2"
)

func Getusers(c *fiber.Ctx) error {
	var users []models.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(users)
}

func Getuser(c *fiber.Ctx) error {
	var user models.User
	id := c.QueryInt("id")
	err := database.DB.First(&user, id).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}
