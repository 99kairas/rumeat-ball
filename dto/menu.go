package dto

import (
	"rumeat-ball/models"

	"github.com/google/uuid"
)

type CreateMenuRequest struct {
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Image       string  `json:"image" form:"image"`
	Price       float64 `json:"price" form:"price"`
}

type CreateMenuResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
}

type UpdateMenuRequest struct {
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Image       string  `json:"image" form:"image"`
	Price       float64 `json:"price" form:"price"`
}

type UpdateMenuResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
}

type GetMenuResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
}

func ConvertToGetMenuResponse(menu models.Menu) GetMenuResponse {
	return GetMenuResponse{
		ID:          menu.ID,
		Name:        menu.Name,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
		Status:      menu.Status,
	}
}

func ConvertToUpdateMenuModel(menu UpdateMenuRequest) models.Menu {
	return models.Menu{
		Name:        menu.Name,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
	}
}

func ConvertToUpdateMenuResponse(menu models.Menu, id uuid.UUID) UpdateMenuResponse {
	return UpdateMenuResponse{
		ID:          id,
		Name:        menu.Name,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
	}
}

func ConvertToCreateMenuModel(menu CreateMenuRequest) models.Menu {
	return models.Menu{
		ID:          uuid.New(),
		Name:        menu.Name,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
		Status:      "available",
	}
}

func ConvertToCreateMenuResponse(menu models.Menu) CreateMenuResponse {
	return CreateMenuResponse{
		ID:          menu.ID,
		Name:        menu.Name,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
	}
}
