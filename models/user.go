package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `json:"id" form:"id" gorm:"type:varchar(100)"`
	Email        string    `json:"email" form:"email" gorm:"type:varchar(50);not null"`
	Name         string    `json:"name" form:"name" gorm:"type:varchar(50);not null"`
	Address      string    `json:"address" form:"address" gorm:"type:text"`
	Phone        string    `json:"phone" form:"phone" gorm:"type:varchar(15)"`
	Password     string    `json:"password" form:"password" gorm:"type:varchar(255)"`
	ProfileImage string    `json:"profile_image" form:"profile_image" gorm:"type:varchar(255)"`
	Status       string    `gorm:"type:enum('verified', 'unverified');default:unverified" json:"status" form:"status"`
	OTP          string    `json:"otp" form:"otp" gorm:"type:varchar(6);default:null"`
	Role         string    `gorm:"type:enum('user', 'admin');default:user" json:"role" form:"role"`
}
