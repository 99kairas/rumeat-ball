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
	ID            uuid.UUID                   `json:"id"`
	Name          string                      `json:"name"`
	Description   string                      `json:"description"`
	Image         string                      `json:"image"`
	Price         float64                     `json:"price"`
	Status        string                      `json:"status"`
	CategoryID    uuid.UUID                   `json:"category_id"`
	CommentCount  int                         `json:"comment_count"`
	AverageRating float64                     `json:"average_rating"`
	Comments      []RatingResponseDetailsMenu `json:"comments"` // Tambahkan field ini
}

type GetAllMenuResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Image         string    `json:"image"`
	Price         float64   `json:"price"`
	Status        string    `json:"status"`
	CategoryID    uuid.UUID `json:"category_id"`
	CommentCount  int       `json:"comment_count"`
	AverageRating float64   `json:"average_rating"`
}

func ConvertToGetMenuResponse(menu models.Menu, commentCount int, averageRating float64) GetMenuResponse {
	return GetMenuResponse{
		ID:            menu.ID,
		Name:          menu.Name,
		Description:   menu.Description,
		Image:         menu.Image,
		Price:         menu.Price,
		Status:        menu.Status,
		CategoryID:    menu.CategoryID,
		CommentCount:  commentCount,
		AverageRating: averageRating,
		Comments:      nil, // Atur ini saat membuat response di controller
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

func ConvertToGetAllMenuResponse(menus []models.Menu, commentCounts map[uuid.UUID]int, averageRatings map[uuid.UUID]float64) []GetAllMenuResponse {
	var response []GetAllMenuResponse
	for _, menu := range menus {
		response = append(response, GetAllMenuResponse{
			ID:            menu.ID,
			Name:          menu.Name,
			Description:   menu.Description,
			Image:         menu.Image,
			Price:         menu.Price,
			Status:        menu.Status,
			CategoryID:    menu.CategoryID,
			CommentCount:  commentCounts[menu.ID],
			AverageRating: averageRatings[menu.ID],
		})
	}
	return response
}
