package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `json:"id" form:"id"`
	Email     string    `json:"email" form:"email"`
	Address   string    `json:"address" form:"address"`
	Phone     string    `json:"phone" form:"phone"`
	Passsword string    `json:"password" form:"password"`
	Status    string    `gorm:"type:enum('verified', 'unverified');default:unverified" json:"status" form:"status"`
}
