package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID    uuid.UUID `json:"id" form:"id" gorm:"type:varchar(100)"`
	Name  string    `json:"name" form:"name" gorm:"type:varchar(100);not null"`
	Menus []Menu    `gorm:"foreignKey:CategoryID"`
}
