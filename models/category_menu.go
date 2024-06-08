package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID    uuid.UUID `json:"id" form:"id"`
	Name  string    `json:"name" form:"name"`
	Menus []Menu    `gorm:"foreignKey:CategoryID"`
}
