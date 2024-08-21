package routes

import (
	"todolist-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func Initroutes(app *fiber.App) {

	app.Get("/users", handlers.Getusers)
	app.Get("/user", handlers.Getuser)
	app.Post("/user/login", handlers.LoginUser)
	app.Post("/user/register", handlers.CreateUser)
	app.Get("/user/tasks", handlers.Gettasks)
	app.Post("/user/tasks", handlers.CreateTask)
	app.Put("/user/tasks", handlers.UpdateTask)
}
