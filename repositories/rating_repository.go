package repositories

import (
	"rumeat-ball/database"
	"rumeat-ball/models"

	"github.com/google/uuid"
)

func CreateRating(data models.Rating) (models.Rating, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return models.Rating{}, tx.Error
	}
	return data, nil
}

func GetRating() ([]models.Rating, error) {
	var data []models.Rating
	tx := database.DB.Find(&data)
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

func UpdateRating(data models.Rating) (models.Rating, error) {
	tx := database.DB.Save(&data)
	if tx.Error != nil {
		return models.Rating{}, tx.Error
	}
	return data, nil
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
