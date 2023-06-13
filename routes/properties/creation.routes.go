package properties

import (
	"github.com/gofiber/fiber/v2"
	crud "github.com/wambua-iv/occupyBE-go/CRUD"
	"github.com/wambua-iv/occupyBE-go/middlewares"
	"github.com/wambua-iv/occupyBE-go/models"
	"github.com/wambua-iv/occupyBE-go/utils"
)

type propertyCreationSerializer struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Type        string `json:"type" validate:"required"`
}

func SetUpPropertyRoutes(app *fiber.App) {

	app.Post("occupy/create_listing", createListing)
	// app.Post("occupy/users/registe", registerUser)
	app.Get("occupy/get_properties", getAllProperties)
	app.Get("occupy/testing_cookie", testCookie)
}

func createListing(ctx *fiber.Ctx) error {
	ctx.Accepts("applicattion/json")
	var body propertyCreationSerializer
	var property models.Property

	keys := middlewares.GetKeys(ctx)
	if len(keys) < 1 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "user not authenticated",
		})
	}

	//confirmation that property owner is the authenticated user
	propertyOwner, err := crud.GetUser(keys[0])
	if err != nil || propertyOwner.Email != keys[0] {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Action is authorized",
		})
	}

	//listing creation data validation
	ctx.BodyParser(&body)
	if err := utils.ValidateStruct(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	//register property
	property.Owner = propertyOwner.ID
	ctx.BodyParser(&property)

	err = crud.CreatePropertyListing(property)
	if err != nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "error in registering user" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON("Properties created")

}

func getAllProperties(ctx *fiber.Ctx) error{
	properties, err := crud.ViewAllProperties()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "you are not authorized to access this route" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(properties)
}


func testCookie(ctx *fiber.Ctx) error {

	keys := middlewares.GetKeys(ctx)
	if keys != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "no cookies" + keys[0],
		})
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Incorrect password provided",
	})
}
