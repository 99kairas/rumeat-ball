package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uuid.UUID `json:"id" form:"id"`
	Date         time.Time `json:"date" form:"date"`
	Status       string    `gorm:"type:enum('processed', 'successed', 'cancelled')" json:"status" form:"status"`
	Total        float64   `json:"total" form:"total"`
	DetailOrders []DetailOrder
}
