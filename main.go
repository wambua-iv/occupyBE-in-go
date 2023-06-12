package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/wambua-iv/occupyBE-go/database"
	"github.com/wambua-iv/occupyBE-go/routes"
)

func register(c *fiber.Ctx) error {
	return c.SendString("Welcome to occupy")
}

func main() {
	occupy := fiber.New()
	occupy.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	database.ConnectDB()
	routes.SetupRoutes(occupy)

	occupy.Get("register", register)
	log.Fatal(occupy.Listen(":3080"))
}
