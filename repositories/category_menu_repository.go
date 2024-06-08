package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/models"

	"github.com/google/uuid"
)

func CreateCategory(category models.Category) (models.Category, error) {
	err := database.DB.Create(&category).Error
	return category, err
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

func UpdateCategory(id uuid.UUID, category models.Category) (models.Category, error) {
	err := database.DB.Model(&category).Where("id = ?", id).Updates(&category).Error
	return category, err
}

func DeleteCategory(id uuid.UUID) error {
	err := database.DB.Delete(&models.Category{}, "id = ?", id).Error
	return err
}
