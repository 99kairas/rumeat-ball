package dto

import (
	"rumeat-ball/models"

	"github.com/google/uuid"
)

type CreateCategoryRequest struct {
	Name string `json:"name" form:"name"`
}

type CreateCategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" form:"name"`
}

type UpdateCategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetCategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func ConvertToCreateCategoryModel(req CreateCategoryRequest) models.Category {
	return models.Category{
		ID:   uuid.New(),
		Name: req.Name,
	}
}

func ConvertToCreateCategoryResponse(category models.Category) CreateCategoryResponse {
	return CreateCategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ConvertToUpdateCategoryModel(req UpdateCategoryRequest) models.Category {
	return models.Category{
		Name: req.Name,
	}
}

func ConvertToUpdateCategoryResponse(category models.Category, id uuid.UUID) UpdateCategoryResponse {
	return UpdateCategoryResponse{
		ID:   id,
		Name: category.Name,
	}
}

func ConvertToGetCategoryResponse(category models.Category) GetCategoryResponse {
	return GetCategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

