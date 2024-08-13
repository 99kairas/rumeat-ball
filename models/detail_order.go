package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DetailOrder struct {
	gorm.Model
	ID         uuid.UUID `json:"id" form:"id" gorm:"type:varchar(100)"`
	OrderID    string    `gorm:"index" json:"order_id" form:"order_id"`
	MenuID     uuid.UUID `json:"menu_id" form:"menu_id"`
	Quantity   int       `json:"quantity" form:"quantity" gorm:"type:integer"`
	TotalPrice float64   `json:"total_price" form:"total_price" gorm:"type:decimal(7,2)"`
}
