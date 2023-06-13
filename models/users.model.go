package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" gorm:"unique" validate:"required,email"`
	Hash      string `json:"password" validate:"required,min=8"`
}

// generate UUID for the model
func (user *User) BeforeCreate(scope *gorm.DB) error {
	id, err := uuid.NewRandom()
	user.ID = id
	return err
}
