package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID         uuid.UUID `json:"id" form:"id"`
	OrderID    uuid.UUID `gorm:"index" json:"order_id" form:"order_id"`
	TotalPrice float64   `json:"total_price" form:"total_price"`
	Status     string    `gorm:"type:enum('successed', 'failed')" json:"status" form:"status"`
	UserID     uuid.UUID `gorm:"index" json:"user_id" form:"user_id"`
	User       User      `gorm:"foreignKey:UserID"`
}
