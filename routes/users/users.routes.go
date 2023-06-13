package users

import (
	"github.com/gofiber/fiber/v2"
	crud "github.com/wambua-iv/occupyBE-go/crud"
	"github.com/wambua-iv/occupyBE-go/middlewares"
	"github.com/wambua-iv/occupyBE-go/models"
	"github.com/wambua-iv/occupyBE-go/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationSerilizer struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Hash      string `json:"password" validate:"required,min=8"`
}

type UserLoginSerilizer struct {
	Email string `json:"email" validate:"required,email"`
	Hash  string `json:"password" validate:"required,min=8"`
}


func SetUpUserRoutes(app *fiber.App) {

	app.Get("occupy/users/all_users", getUsers)
	app.Post("occupy/users/register", registerUser)
	app.Post("occupy/users/login", userLogin)
}

func getUsers(c *fiber.Ctx) error {
	users, err := crud.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "you are not authorized to access this route" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func registerUser(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	var valid UserLoginSerilizer
	var user models.User

	//user input validation
	ctx.BodyParser(&valid)
	if err := utils.ValidateStruct(valid); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	//user password encryption
	hash := []byte(user.Hash)
	hashedPassword, err := bcrypt.GenerateFromPassword(hash, bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	user.Hash = string(hashedPassword)

	//create user on database
	ctx.BodyParser(&user)
	err = crud.Createuser(user)
	if err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "error in registering user" + err.Error(),
		})
	}
	return ctx.SendString("User Successfully created")
}

func userLogin(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	var user UserLoginSerilizer

	ctx.BodyParser(&user)
	if err := utils.ValidateStruct(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	existingUser, err := crud.GetUser(user.Email)
	if err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "error in registering user" + err.Error(),
		})
	}

	// Comparing the password with the hash
	password := []byte(user.Hash)
	hashedPassword := []byte(existingUser.Hash)
	if err = bcrypt.CompareHashAndPassword(hashedPassword, password); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Incorrect password provided",
		})
	}
	
	//create user session
	
	currentSession := middlewares.SetSession(ctx, existingUser.Email)
	if currentSession == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "error in registering user",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(currentSession)
}


