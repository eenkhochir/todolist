package handlers

import (
	"todolist-backend/database"
	"todolist-backend/models"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {
	loginData := new(models.User)
	var user models.User
	c.BodyParser(loginData)
	database.DB.Where("username = ? OR email = ? OR phonenumber = ?", loginData.Username, loginData.Email, loginData.Phonenumber).First(&user)

	if user.ID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Хэрэглэгч олдсонгүй."})
	}

	if user.Password != loginData.Password {
		return c.Status(400).JSON(fiber.Map{"error": "Нууц үг буруу."})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Амжилттай нэвтэрлээ."})
}
