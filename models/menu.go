package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	ID          uuid.UUID `json:"id" form:"id" gorm:"type:varchar(100)"`
	Name        string    `json:"name" form:"name" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" form:"description" gorm:"type:text"`
	Image       string    `json:"image" form:"image" gorm:"type:varchar(255)"`
	Price       float64   `json:"price" form:"price" gorm:"type:decimal(7,2)"`
	Status      string    `gorm:"type:enum('available', 'unavailable');default:available" json:"status" form:"status"`
	CategoryID  uuid.UUID `gorm:"type:varchar(100);not null;index" json:"category_id" form:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
}
