package repositories

import (
	"errors"
	"rumeat-ball/database"
	"rumeat-ball/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateRating(data models.Rating) (models.Rating, error) {
	var existingRating models.Rating
	tx := database.DB.Where("user_id = ? AND menu_id = ?", data.UserID, data.MenuID).First(&existingRating)

	if tx.Error == nil {
		// Rating already exists
		return models.Rating{}, errors.New("rating already exists")
	} else if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return models.Rating{}, tx.Error
	}

	// Rating does not exist, create a new one
	data.ID = uuid.New()
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	tx = database.DB.Create(&data)
	if tx.Error != nil {
		return models.Rating{}, tx.Error
	}

	return data, nil
}

func GetAllRatings(userID uuid.UUID) ([]models.Rating, error) {
	var data []models.Rating
	tx := database.DB.Where("user_id = ?", userID).Preload("Menu").Preload("User").Find(&data)

	if tx.Error != nil {
		return []models.Rating{}, tx.Error
	}

	return data, nil
}

func GetRatingByID(id uuid.UUID) (models.Rating, error) {
	var data models.Rating
	tx := database.DB.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return models.Rating{}, tx.Error
	}
	return data, nil
}

func UpdateRating(data models.Rating, ratingID uuid.UUID) (models.Rating, error) {
	var existingRating models.Rating
	tx := database.DB.Where("id = ?", ratingID).Preload("Menu").Preload("User").First(&existingRating)

	if tx.Error != nil {
		return models.Rating{}, tx.Error
	}

	// Update the rating
	existingRating.Rating = data.Rating

	// Check if the new comment is not empty before updating
	if data.Comment != "" {
		existingRating.Comment = data.Comment
	}

	existingRating.UpdatedAt = time.Now()
	tx = database.DB.Save(&existingRating)
	if tx.Error != nil {
		return models.Rating{}, tx.Error
	}

	return existingRating, nil
}

func DeleteRating(id uuid.UUID) error {
	tx := database.DB.Where("id = ?", id).Delete(&models.Rating{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetRatingByUserID(id uuid.UUID) ([]models.Rating, error) {
	var data []models.Rating
	tx := database.DB.Where("user_id = ?", id).Find(&data)
	if tx.Error != nil {
		return []models.Rating{}, tx.Error
	}
	return data, nil
}

func GetRatingByMenuID(id uuid.UUID) ([]models.Rating, error) {
	var data []models.Rating
	tx := database.DB.Where("menu_id = ?", id).Find(&data)
	if tx.Error != nil {
		return []models.Rating{}, tx.Error
	}
	return data, nil
}

func GetRatingByOrderID(id uuid.UUID) ([]models.Rating, error) {
	var data []models.Rating
	tx := database.DB.Where("order_id = ?", id).Find(&data)
	if tx.Error != nil {
		return []models.Rating{}, tx.Error
	}
	return data, nil
}

func GetRatingByUserIDAndMenuID(userID, menuID uuid.UUID) (models.Rating, error) {
	var data models.Rating
	tx := database.DB.Where("user_id = ? AND menu_id = ?", userID, menuID).First(&data)
	if tx.Error != nil {
		return models.Rating{}, tx.Error
	}
	return data, nil
}

func GetRatingByUserIDAndOrderID(userID, orderID uuid.UUID) (models.Rating, error) {
	var data models.Rating
	tx := database.DB.Where("user_id = ? AND order_id = ?", userID, orderID).First(&data)
	if tx.Error != nil {
		return models.Rating{}, tx.Error
	}
	return data, nil
}
