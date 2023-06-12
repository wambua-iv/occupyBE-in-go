package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wambua-iv/occupyBE-go/routes/users"
)

func SetupRoutes(app *fiber.App) {
	users.SetUpUserRoutes(app)
}
