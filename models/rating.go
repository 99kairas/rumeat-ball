package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	ID      uuid.UUID `json:"id" form:"id" gorm:"type:varchar(100)"`
	MenuID  uuid.UUID `gorm:"index" json:"menu_id" form:"menu_id"`
	Menu    Menu      `gorm:"foreignKey:MenuID"`
	UserID  uuid.UUID `gorm:"index" json:"user_id" form:"user_id"`
	User    User      `gorm:"foreignKey:UserID"`
	Rating  float64   `json:"rating" form:"rating" gorm:"type:decimal(7,2)"`
	Comment string    `json:"comment" form:"comment" gorm:"type:text"`
	Date    time.Time `json:"date" form:"date"`
}
