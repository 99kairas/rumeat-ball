package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	ID      uuid.UUID `json:"id" form:"id"`
	MenuID  uuid.UUID `gorm:"index" json:"menu_id" form:"menu_id"`
	UserID  uuid.UUID `gorm:"index" json:"user_id" form:"user_id"`
	Rating  int       `json:"rating" form:"rating"`
	Comment string    `json:"comment" form:"comment" gorm:"type:text"`
	Date    time.Time `json:"date" form:"date"`
}
