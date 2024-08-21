package handlers

import (
	"todolist-backend/database"
	"todolist-backend/models"

	"github.com/gofiber/fiber/v2"
)

func Gettasks(c *fiber.Ctx) error {
	userID := c.QueryInt("id")

	// fmt.Println("asuudalgui")
	var tasks []models.Todo
	database.DB.Where("userid = ?", userID).Find(&tasks)
	return c.Status(200).JSON(tasks)
	// var user models.User
	// database.DB.Preload("Tasks").First(&user, userID)
	// var tasks []models.Todo
	// for _, task := range user.Tasks {
	// 	// fmt.Printf("Task: %s, Description: %s\n", task.Taskname, task.Description)
	// 	tasks = append(tasks, task)
	// }
	// return c.Status(200).JSON(tasks)
}
func CreateTask(c *fiber.Ctx) error {
	var task models.Todo

	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "wrong input",
		})
	}

	db := database.DB
	if err := db.Create(&task).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.QueryInt("id")
	var task models.Todo
	err := database.DB.First(&task, id).Error
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"messsage": err.Error(),
		})
	}
	err = c.BodyParser(&task)
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err = database.DB.Save(&task).Error
	if err != nil {
		c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(task)
}
