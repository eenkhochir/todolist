package main

import (
	"log"
	"todolist-backend/database"
	"todolist-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Initdatabase()

	app := fiber.New()

	routes.Initroutes(app)

	log.Fatal(app.Listen(":3000"))
}
