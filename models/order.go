package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           string    `json:"id" form:"id"`
	UserID       uuid.UUID `gorm:"index" json:"user_id" form:"user_id"`
	User         User      `gorm:"foreignKey:UserID"`
	Date         time.Time `json:"date" form:"date"`
	Status       string    `gorm:"type:enum('processed', 'successed', 'cancelled', 'cart')" json:"status" form:"status"`
	Total        float64   `json:"total" form:"total"`
	DetailOrders []DetailOrder
}
