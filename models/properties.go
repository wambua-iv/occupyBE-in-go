package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Property struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt   time.Time
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	Address     string    `json:"address" validate:"required"`
	Type        string    `json:"type" validate:"required"`
	Owner       uuid.UUID `json:"owner" gorm:"type:uuid;foreignKey"`
}

func (property *Property) BeforeCreate(scope *gorm.DB) error {
	id, err := uuid.NewRandom()
	property.ID = id
	return err
}
