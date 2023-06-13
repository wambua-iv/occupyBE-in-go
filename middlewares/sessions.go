package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v2"
)


var storage = postgres.New(postgres.Config{
	Host:       "127.0.0.1",
	Port:       5050,
	Database:   "rainbow_database",
	Table:      "fiber_storage",
	SSLMode:    "disable",
	Reset:      false,
	GCInterval: 10 * time.Second,
	Username:   "unicorn_user",
	Password:   "magical_password",
})

var Store = session.New(session.Config{
	Storage: storage})

func GetKeys(ctx *fiber.Ctx) []string {
	currentSession, err := Store.Get(ctx)
	if err != nil {
		ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "Incorrect password provided",
		})
		
		return currentSession.Keys()
	}

	return currentSession.Keys()
	
}

func SetSession(ctx *fiber.Ctx, email string) []string {
	currentSession, err := Store.Get(ctx)
	if err != nil {
		panic(err)
	}
	currentSession.Set(email, email)
	err = currentSession.Save()
	if err != nil {
		ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "error in registering user" + err.Error(),
		})
		return []string{"session not created"}
	}
	return currentSession.Keys()
}

// func testCookie(ctx *fiber.Ctx) error {
// 	ctx.Accepts("application/json")
// 	currentSession, err := Store.Get(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	type User struct {
// 		Email string `json:"email" validate:"required,email"`
// 	}
// 	var user User
// 	ctx.BodyParser(&user)
// 	keys := currentSession.Keys()
// 	if len(keys) > 0 {
// 		return ctx.Status(fiber.StatusOK).JSON(keys[0])
// 	}
// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 		"message": "Incorrect password provided",
// 	})
// }
