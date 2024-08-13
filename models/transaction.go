package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID         uuid.UUID `json:"id" form:"id" gorm:"type:varchar(100)"`
	OrderID    string    `gorm:"index;type:varchar(10)" json:"order_id" form:"order_id"`
	TotalPrice float64   `json:"total_price" form:"total_price" gorm:"type:decimal(7,2)"`
	PaymentURL string    `json:"payment_url" form:"payment_url" gorm:"type:text"`
	Status     string    `gorm:"type:enum('successed', 'failed', 'pending')" json:"status" form:"status"`
	UserID     uuid.UUID `gorm:"index" json:"user_id" form:"user_id"`
	User       User      `gorm:"foreignKey:UserID"`
}
