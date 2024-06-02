package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	ID          uuid.UUID `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description" gorm:"type:text"`
	Image       string    `json:"image" form:"image" `
	Price       float64   `json:"price" form:"price"`
	Status      string    `gorm:"type:enum('available', 'unavailable');default:available" json:"status" form:"status"`
}
