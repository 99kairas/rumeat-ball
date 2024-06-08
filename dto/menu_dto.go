package dto

import (
	"rumeat-ball/models"

	"github.com/google/uuid"
)

type CreateMenuRequest struct {
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Image       string    `json:"image" form:"image"`
	Price       float64   `json:"price" form:"price"`
	CategoryID  uuid.UUID `json:"category_id" form:"category_id"`
}

type CreateMenuResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	CategoryID  uuid.UUID `json:"category_id"`
}

type UpdateMenuRequest struct {
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Image       string    `json:"image" form:"image"`
	Price       float64   `json:"price" form:"price"`
	CategoryID  uuid.UUID `json:"category_id" form:"category_id"`
}

type UpdateMenuResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	CategoryID  uuid.UUID `json:"category_id"`
}

type GetMenuResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	CategoryID  uuid.UUID `json:"category_id"`
}

func ConvertToGetMenuResponse(menu models.Menu) GetMenuResponse {
	return GetMenuResponse{
		ID:          menu.ID,
		Name:        menu.Name,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
		Status:      menu.Status,
		CategoryID:  menu.CategoryID,
	}
}

func ConvertToUpdateMenuModel(menu UpdateMenuRequest) models.Menu {
	return models.Menu{
		Name:        menu.Name,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
		CategoryID:  menu.CategoryID,
	}
}

func ConvertToUpdateMenuResponse(menu models.Menu, id uuid.UUID) UpdateMenuResponse {
	return UpdateMenuResponse{
		ID:          id,
		Name:        menu.Name,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
		CategoryID:  menu.CategoryID,
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
		CategoryID:  menu.CategoryID,
	}
}

func ConvertToCreateMenuResponse(menu models.Menu) CreateMenuResponse {
	return CreateMenuResponse{
		ID:          menu.ID,
		Name:        menu.Name,
		Description: menu.Description,
		Image:       menu.Image,
		Price:       menu.Price,
		CategoryID:  menu.CategoryID,
	}
}
