package dto

import {
	"github.com/google/uuid"
	"time"
	}

type OrderRequest struct {
	UserID string `json:"user_id" form:"user_id"`
	Date   time.Time `json:"date" form:"date"`
}

type OrderResponse struct {
	ID     uuid.UUID `json:"id" form:"id"`
	Date   time.Time `json:"date" form:"date"`
	Status string    `json:"status" form:"status"`
}
