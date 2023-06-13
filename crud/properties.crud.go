package crud

import (
	"github.com/wambua-iv/occupyBE-go/database"
	"github.com/wambua-iv/occupyBE-go/models"
)

func ViewAllProperties()([]models.Property, error) {
	var properties []models.Property
	transaction := database.Database.Db.Find(&properties)

	if transaction.Error != nil {
		return []models.Property{}, transaction.Error
	}
	return properties, nil
}
func CreatePropertyListing(property models.Property) error {
	transaction := database.Database.Db.Create(&property)
	return transaction.Error
}