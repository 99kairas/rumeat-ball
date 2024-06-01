package dto

import (
	"rumeat-ball/models"

	"github.com/google/uuid"
)

type CreateMenuRequest struct {
	Description string  `json:"description" form:"description"`
	Image       string  `json:"image" form:"image"`
	Price       float64 `json:"price" form:"price"`
}

type CreateMenuResponse struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
}

func ConvertToCreateMenuModel(menu CreateMenuRequest) models.Menu {
	return models.Menu{
		ID:          uuid.New(),
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
		Status:      "available",
	}
}

func ConvertToCreateMenuResponse(menu models.Menu) CreateMenuResponse {
	return CreateMenuResponse{
		ID:          menu.ID,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
	}
}
