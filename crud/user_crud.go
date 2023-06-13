package crud

import (
	"github.com/wambua-iv/occupyBE-go/database"
	"github.com/wambua-iv/occupyBE-go/models"
)

type ContextUser models.User

func (user *ContextUser) ShowProfile(userData models.User) {

}

func Createuser(user models.User) error {
	transaction := database.Database.Db.Create(&user)
	return transaction.Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	transaction := database.Database.Db.Find(&users)

	if transaction.Error != nil {
		return []models.User{}, transaction.Error
	}
	return users, nil
}

func GetUser(email string) (models.User, error) {
	var user models.User

	transaction := database.Database.Db.Where("Email = ?", email).First(&user)
	if transaction.Error != nil {
		return models.User{}, transaction.Error
	}
	return user, nil
}

func UpdateUser(user models.User) error {
	transaction := database.Database.Db.Save(&user)
	return transaction.Error
}

func DeleteUser(id uint64) error {
	transaction := database.Database.Db.Unscoped().Delete(&models.User{})
	return transaction.Error
}
